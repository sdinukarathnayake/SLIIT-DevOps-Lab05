package com.sliit.microservices.apigateway.config;

import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.cors.CorsConfiguration;
import org.springframework.web.cors.reactive.CorsWebFilter;
import org.springframework.web.cors.reactive.UrlBasedCorsConfigurationSource;

import java.util.Arrays;

@Configuration
public class GatewayConfig {

    @Bean
    public RouteLocator customRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
                // Item Service Routes with detailed path matching
                .route("item-service-root", r -> r
                    .path("/api/items")
                    .filters(f -> f.stripPrefix(1))
                    .uri("http://item-service:8081"))
                .route("item-service-all", r -> r
                    .path("/api/items/**")
                    .filters(f -> f.stripPrefix(1))
                    .uri("http://item-service:8081"))
                
                // Order Service Routes
                .route("order-service-root", r -> r
                    .path("/api/orders")
                    .filters(f -> f.stripPrefix(1))
                    .uri("http://order-service:8082"))
                .route("order-service-all", r -> r
                    .path("/api/orders/**")
                    .filters(f -> f.stripPrefix(1))
                    .uri("http://order-service:8082"))
                
                // Payment Service Routes
                .route("payment-service-root", r -> r
                    .path("/api/payments")
                    .filters(f -> f.stripPrefix(1))
                    .uri("http://payment-service:8083"))
                .route("payment-service-all", r -> r
                    .path("/api/payments/**")
                    .filters(f -> f.stripPrefix(1))
                    .uri("http://payment-service:8083"))
                
                // Health check routes for each service
                .route("item-health", r -> r
                    .path("/api/items/health")
                    .filters(f -> f.stripPrefix(1).rewritePath("/health", "/"))
                    .uri("http://item-service:8081"))
                .route("order-health", r -> r
                    .path("/api/orders/health")
                    .filters(f -> f.stripPrefix(1))
                    .uri("http://order-service:8082"))
                .route("payment-health", r -> r
                    .path("/api/payments/health")
                    .filters(f -> f.stripPrefix(1).rewritePath("/health", "/docs"))
                    .uri("http://payment-service:8083"))
                
                .build();
    }

    @Bean
    public CorsWebFilter corsWebFilter() {
        CorsConfiguration corsConfiguration = new CorsConfiguration();
        corsConfiguration.setAllowedOriginPatterns(Arrays.asList("*"));
        corsConfiguration.setAllowedMethods(Arrays.asList("GET", "POST", "PUT", "DELETE", "OPTIONS"));
        corsConfiguration.setAllowedHeaders(Arrays.asList("*"));
        corsConfiguration.setAllowCredentials(false);

        UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
        source.registerCorsConfiguration("/**", corsConfiguration);

        return new CorsWebFilter(source);
    }
}
