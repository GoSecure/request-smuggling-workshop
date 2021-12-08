package net.gosecure.websocket.config;

import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.util.Map;

import static org.springframework.http.MediaType.TEXT_PLAIN;

@RestController
public class HealthCheckController {

    @RequestMapping("/health-check")
    public ResponseEntity<String> health(@RequestParam("url") String url,
                                         @RequestParam(value = "code", required = false) Integer statusCodeInput)
            throws IOException {

//        String url = body.get("url");
//        Integer statusCodeInput = Integer.parseInt(body.get("code"));

        OkHttpClient client = new OkHttpClient();

        Request request = new Request.Builder()
                .url(url)
                .build();

        if(url.equals("http://store.initech.com")) {
            statusCodeInput = 200;
        }
        else if(url.equals("http://tickets.initech.com")) {
            statusCodeInput = 200;
        }
        else if(url.equals("http://blog.initech.com")) {
            statusCodeInput = 500;
        }


        int statusCode;


        if(statusCodeInput != null) {
            statusCode = statusCodeInput;
        }
        else {
            try (Response response = client.newCall(request).execute()) {
                statusCode = response.code();
            }
        }

        return ResponseEntity
                .status(statusCode)
                .contentType(TEXT_PLAIN)
                .body("");
    }
}
