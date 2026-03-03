package com.sliit.microservices.apigateway.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.HashMap;
import java.util.Map;

@RestController
@RequestMapping("/gateway")
public class GatewayController {

    @Autowired
    private RouteLocator routeLocator;

    @GetMapping("/")
    public Mono<Map<String, Object>> getGatewayInfo() {
        Map<String, Object> info = new HashMap<>();
        info.put("name", "SLIIT Microservices API Gateway");
        info.put("version", "1.0.0");
        info.put("description", "Central API Gateway for routing requests to microservices");
        
        Map<String, String> services = new HashMap<>();
        services.put("Item Service", "http://localhost:8081 -> /api/items/**");
        services.put("Order Service", "http://localhost:8082 -> /api/orders/**");
        services.put("Payment Service", "http://localhost:8083 -> /api/payments/**");
        
        info.put("services", services);
        info.put("status", "running");
        
        return Mono.just(info);
    }

    @GetMapping("/routes")
    public Flux<Map<String, String>> getRoutes() {
        return routeLocator.getRoutes()
                .map(route -> {
                    Map<String, String> routeInfo = new HashMap<>();
                    routeInfo.put("id", route.getId());
                    routeInfo.put("uri", route.getUri().toString());
                    routeInfo.put("predicates", route.getPredicate().toString());
                    return routeInfo;
                });
    }

    @GetMapping("/health")
    public Mono<Map<String, String>> healthCheck() {
        Map<String, String> health = new HashMap<>();
        health.put("status", "UP");
        health.put("service", "api-gateway");
        health.put("timestamp", java.time.Instant.now().toString());
        return Mono.just(health);
    }
}
