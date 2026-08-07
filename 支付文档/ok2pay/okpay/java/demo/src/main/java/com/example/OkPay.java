package com.example;

import org.json.JSONArray;
import org.json.JSONObject;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URL;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.util.*;

public class OkPay {
    protected String id;
    protected String token;
    protected String api_url_payLink;
    protected String api_url_transfer;
    protected String api_url_TransactionHistory;
    protected String api_url_checkTransferByTxid;
    protected String api_url_checkTransfer;
    protected String url;

    public OkPay(String id, String token) {
        this.id = id;
        this.token = token;

        String api_url = "https://api.okaypay.me/shop/";
        this.api_url_payLink = api_url + "payLink";
        this.api_url_transfer = api_url + "transfer";
        this.api_url_TransactionHistory = api_url + "TransactionHistory";
        this.api_url_checkTransferByTxid = api_url + "checkTransferByTxid";
        this.api_url_checkTransfer = api_url + "checkTransfer";
    }

    public JSONObject checkTransfer(Map<String, String> data) {
        this.url = this.api_url_checkTransfer;
        return this.post(data);
    }

    public JSONObject checkTransferByTxid(Map<String, String> data) {
        this.url = this.api_url_checkTransferByTxid;
        return this.post(data);
    }

    public JSONObject payLink(Map<String, String> data) {
        this.url = this.api_url_payLink;
        return this.post(data);
    }

    public JSONObject transfer(Map<String, String> data) {
        this.url = this.api_url_transfer;
        return this.post(data);
    }

    public JSONObject shop_transaction_history(Map<String, String> data) {
        this.url = this.api_url_TransactionHistory;
        return this.post(data);
    }

    public void notify(Map<String, String> data) {
        if (this.checkSign(new JSONObject(data))) {
            System.out.println("验证成功");
            if (data.get("status").equals("success") && data.containsKey("code") && data.get("code").equals("10000")) {
                // 数据正常
            } else {
                // 数据不正常
            }
        } else {
            System.out.println("验证失败");
        }
    }

    public Map<String, String> sign(Map<String, String> data) {
        data.put("id", this.id);
        TreeMap<String, String> sortedData = new TreeMap<>(data);
        StringBuilder query = new StringBuilder();
        for (Map.Entry<String, String> entry : sortedData.entrySet()) {
            if (entry.getValue() != null && !entry.getValue().isEmpty()) {
                if (query.length() > 0) {
                    query.append("&");
                }
                query.append(entry.getKey()).append("=").append(entry.getValue());
            }
        }
        query.append("&token=").append(this.token);
        
        String sign = md5(query.toString()).toUpperCase();
        sortedData.put("sign", sign);
        return sortedData;
    }

    public boolean checkSign(JSONObject json) {
        if (json == null) return false;

        String inSign = json.optString("sign", "");
        if (inSign.isEmpty()) return false;

        Map<String, String> flat = new HashMap<>();
        flattenToBracketMap("", json, flat);
        // flatten 会把 sign 也放进去，这里删掉
        flat.remove("sign");

        String built = buildSignString(flat);
        String expect = md5(built).toUpperCase();

        // 需要的话可打开这两行排查（会打印 token）
        System.out.println(built);
        System.out.println("expect=" + expect + "  actual=" + inSign);

        return inSign.equalsIgnoreCase(expect);
    }

    private String buildSignString(Map<String, String> params) {
        List<Map.Entry<String, String>> entries = new ArrayList<>(params.entrySet());
        entries.sort(signKeyComparator());

        StringBuilder sb = new StringBuilder();
        for (Map.Entry<String, String> e : entries) {
            String k = e.getKey();
            String v = e.getValue();
            if (v == null || v.isEmpty()) continue;

            if (sb.length() > 0) sb.append("&");
            sb.append(k).append("=").append(v);
        }
        sb.append("&token=").append(this.token);
        return sb.toString();
    }

    /**
     * 让拼串更贴近你给的“正确格式”：
     * code -> data[amount],data[unique_id],data[order_id],data[pay_user_id],data[coin],data[status],data[pay_url] -> id -> status -> 其它(字典序)
     */
    private Comparator<Map.Entry<String, String>> signKeyComparator() {
        final Map<String, Integer> topPri = new HashMap<>();
        topPri.put("code", 0);
        topPri.put("id", 2);
        topPri.put("status", 3);

        final Map<String, Integer> dataPri = new HashMap<>();
        dataPri.put("amount", 0);
        dataPri.put("unique_id", 1);
        dataPri.put("order_id", 2);
        dataPri.put("pay_user_id", 3);
        dataPri.put("coin", 4);
        dataPri.put("status", 5);
        dataPri.put("pay_url", 6);

        return (a, b) -> {
            String ka = a.getKey();
            String kb = b.getKey();

            boolean aIsData = ka.startsWith("data[");
            boolean bIsData = kb.startsWith("data[");

            // code 最前
            int pa = topPri.getOrDefault(ka, aIsData ? 1 : 9);
            int pb = topPri.getOrDefault(kb, bIsData ? 1 : 9);
            if (pa != pb) return Integer.compare(pa, pb);

            // data[...] 内按固定优先级
            if (aIsData && bIsData) {
                String da = extractDataField(ka);
                String db = extractDataField(kb);
                int pda = dataPri.getOrDefault(da, 999);
                int pdb = dataPri.getOrDefault(db, 999);
                if (pda != pdb) return Integer.compare(pda, pdb);
                return ka.compareTo(kb);
            }

            // 其它按字典序，保证稳定
            return ka.compareTo(kb);
        };
    }

    private String extractDataField(String k) {
        // data[xxx] 或 data[xxx][yyy]，这里只取第一层字段 xxx
        int start = k.indexOf('[');
        int end = k.indexOf(']', start + 1);
        if (start >= 0 && end > start) return k.substring(start + 1, end);
        return k;
    }

    /**
     * JSONObject -> 扁平化成 a[b][c]=v 形式
     */
    private void flattenToBracketMap(String prefix, Object node, Map<String, String> out) {
        if (node == null) return;

        if (node instanceof JSONObject jo) {
            for (String key : jo.keySet()) {
                Object v = jo.get(key);
                String nextKey = prefix.isEmpty() ? key : (prefix + "[" + key + "]");
                flattenToBracketMap(nextKey, v, out);
            }
            return;
        }

        if (node instanceof JSONArray ja) {
            // 如果未来出现数组，这里用下标展开：arr[0]=...（看文档是否需要）
            for (int i = 0; i < ja.length(); i++) {
                Object v = ja.get(i);
                String nextKey = prefix + "[" + i + "]";
                flattenToBracketMap(nextKey, v, out);
            }
            return;
        }

        // 普通值
        out.put(prefix, String.valueOf(node));
    }

    public JSONObject post(Map<String, String> data) {
        data = this.sign(data);
        System.err.println(data);
        try {
            URL obj = new URL(this.url);
            HttpURLConnection con = (HttpURLConnection) obj.openConnection();
            con.setRequestMethod("POST");
            con.setRequestProperty("User-Agent", "HTTP CLIENT");
            con.setDoOutput(true);
            
            OutputStream os = con.getOutputStream();
            os.write(getPostDataString(data).getBytes());
            os.flush();
            os.close();
            
            int responseCode = con.getResponseCode();
            if (responseCode == HttpURLConnection.HTTP_OK) {
                BufferedReader in = new BufferedReader(new InputStreamReader(con.getInputStream()));
                String inputLine;
                StringBuilder response = new StringBuilder();
                while ((inputLine = in.readLine()) != null) {
                    response.append(inputLine);
                }
                in.close();
                return new JSONObject(response.toString());
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return null;
    }

    private String getPostDataString(Map<String, String> params) throws Exception {
        StringBuilder result = new StringBuilder();
        boolean first = true;
        for (Map.Entry<String, String> entry : params.entrySet()) {
            if (first)
                first = false;
            else
                result.append("&");
            result.append(URLEncoder.encode(entry.getKey(), "UTF-8"));
            result.append("=");
            result.append(URLEncoder.encode(entry.getValue(), "UTF-8"));
        }
        return result.toString();
    }

    private static String md5(String input) {
        try {
            MessageDigest md = MessageDigest.getInstance("MD5");
            byte[] messageDigest = md.digest(input.getBytes(StandardCharsets.UTF_8));
            StringBuilder hexString = new StringBuilder();
            for (byte b : messageDigest) {
                String hex = Integer.toHexString(0xFF & b);
                if (hex.length() == 1) {
                    hexString.append('0');
                }
                hexString.append(hex);
            }
            return hexString.toString();
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}