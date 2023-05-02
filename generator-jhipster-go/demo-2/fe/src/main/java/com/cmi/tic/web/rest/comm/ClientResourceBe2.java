package com.cmi.tic.web.rest.comm;

import com.cmi.tic.config.webClient.AccessToken;
import com.cmi.tic.config.webClient.WebClientConfig;
import java.util.HashMap;
import java.util.Map;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import reactor.core.publisher.Mono;

@RestController
@RequestMapping("/api")
public class ClientResourceBe2 {

    private final Logger log = LoggerFactory.getLogger(ClientResourceBe2.class);

    @Autowired
    private WebClientConfig webClient;

    @Autowired
    private AccessToken oAuthToken;

    private static final String TOKEN_TYPE = "Bearer ";

    public ClientResourceBe2(WebClientConfig webClientConfig, AccessToken oAuth) {
        this.webClient = webClientConfig;
        this.oAuthToken = oAuth;
    }

    @GetMapping("/services/be2")
    public Mono<ResponseEntity<Map<String, String>>> getClient() {
        try {
            Map<String, String> response = new HashMap<>();
            String accessToken = oAuthToken.getToken();
            if (accessToken.isEmpty() || accessToken.isBlank()) {
                log.error("No access token found");
                response.put("server_error", "Please hit the url again!");
                response.put("cause", "NO_ACCESS_TOKEN");
                return Mono.just(ResponseEntity.accepted().body(response));
            }
            return webClient
                .serviceWebClient()
                .build()
                .get()
                .uri("http://be2/api/services/be2")
                .header(HttpHeaders.AUTHORIZATION, TOKEN_TYPE + accessToken)
                .retrieve()
                .bodyToMono(new ParameterizedTypeReference<Map<String, String>>() {})
                .flatMap(hashmap -> {
                    log.info(hashmap.get("server"));
                    response.put("server", hashmap.get("server"));
                    return Mono.just(ResponseEntity.ok(response));
                })
                .onErrorResume(e -> {
                    if (e instanceof Exception) {
                        log.error("Exception due to :- ".concat(e.getMessage()));
                        response.put("server_error", e.getLocalizedMessage());
                    }
                    return Mono.just(ResponseEntity.ok(response));
                });
        } catch (Exception e) {
            log.error(e.getLocalizedMessage());
        }
        return null;
    }
}
