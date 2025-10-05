use axum::{
    Json,
    Router,
    // extract::{Path, State},
    routing::get,
};
use serde::{Deserialize, Serialize};
// use std::sync::Arc;
use utoipa::{OpenApi, ToSchema};
use tokio::net::TcpListener;
use utoipa_swagger_ui::SwaggerUi;
use tower_http::trace::{Trace, TraceLayer};
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};
// use uuid::Uuid;

#[derive(Serialize, Deserialize, ToSchema)]
struct HelloResponse {
    message: String,
}

#[utoipa::path(
    get,
    path = "/",
    responses(
        (status = 200, description = "API is alive")
    )
)]
async fn health() -> &'static str {
    "API is up b"
}

#[utoipa::path(
    get,
    path = "/hello",
    responses(
        (status = 200, description = "Returns a hello message", body = HelloResponse)
    )
)]
async fn hello() -> Json<HelloResponse> {
    Json(HelloResponse {
        message: "Hello from axum".into(),
    })
}

#[derive(OpenApi)]
#[openapi(
    paths(
        hello, 
        health,
    ),
    components(schemas(HelloResponse)),
    tags((name = "Oxide Backend rust edition", description = "I'm too bored"))
)]
struct ApiDoc;

#[tokio::main]
async fn main() {
    tracing_subscriber::registry()
        .with(tracing_subscriber::EnvFilter::new("info,tower_http=debug"))
        .with(tracing_subscriber::fmt::layer())
        .init();

    let app = Router::new()
        .route("/", get(health))
        .route("/hello", get(hello))
        .merge(SwaggerUi::new("/docs").url("/api-docs/openapi.json", ApiDoc::openapi()))
        .layer(TraceLayer::new_for_http());

    println!("Server running at port 8000");

    let listener = TcpListener::bind("0.0.0.0:8000").await.unwrap();

    axum::serve(listener, app).await.unwrap();
}
