package com.example;

import org.json.JSONObject;
import java.util.HashMap;
import java.util.Map;
import java.util.Iterator;



public class Main {
    public static void main(String[] args) {
        
        String id = "";
        String token = "";
        OkPay okPay = new OkPay(id, token);

        // payLink
        Map<String, String> data = new HashMap<>();
        data.put("unique_id", "test7");
        data.put("amount", "30.00");
        data.put("return_url", "https://127.0.0.1/");
        data.put("callback_url", "https://127.0.0.1/callback");
        data.put("coin", "USDT");

        JSONObject response = okPay.payLink(data);
        System.out.println(response);

         // 检查 response 签名
        if (response != null) {
            boolean signOk = okPay.checkSign(response);
            System.out.println("response sign valid: " + signOk);
        }
        
    }


    private static void flattenJsonObject(String prefix, JSONObject jsonObject, Map<String, String> data) {
        Iterator<String> keys = jsonObject.keys();
        while (keys.hasNext()) {
            String key = keys.next();
            Object value = jsonObject.get(key);
            String newKey = prefix.isEmpty() ? key : prefix + "." + key;

            if (value instanceof JSONObject) {
                flattenJsonObject(newKey, (JSONObject) value, data);
            } else {
                data.put(newKey, String.valueOf(value));
            }
        }
    }


    
}