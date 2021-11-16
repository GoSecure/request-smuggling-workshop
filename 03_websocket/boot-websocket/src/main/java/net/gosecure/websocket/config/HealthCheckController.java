package net.gosecure.websocket.config;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import static org.springframework.http.MediaType.TEXT_PLAIN;

@RestController
public class HealthCheckController {

    @RequestMapping("/health-check")
    public ResponseEntity<String> health(@RequestParam("url") String url, @RequestParam("code") int statusCode) {
        return ResponseEntity
                .status(statusCode)
                .contentType(TEXT_PLAIN)
                .body("");
    }
}
