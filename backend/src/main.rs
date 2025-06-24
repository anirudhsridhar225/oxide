use axum::{Router, extract::Path, response::Json, routing::get};
use serde::Serialize;
use tokio::net::TcpListener;
use tower_http::cors::{Any, CorsLayer};

mod utils;

#[derive(Serialize)]
struct Message {
    message: String,
}

// async fn hello() -> Json<Message> {
//     Json(Message {
//         message: "Hello from Axum!!".to_string(),
//     })
// }

async fn greet(Path(name): Path<String>) -> Json<Message> {
    Json(Message {
        message: format!("Hello and welcome to nuxt/axum Mr.{}", name),
    })
}

#[tokio::main]
async fn main() {
    let cors = CorsLayer::new()
        .allow_origin(Any)
        .allow_methods(Any)
        .allow_headers(Any);

    let app = Router::new()
        .route("/api/greet/{name}", get(greet))
        .layer(cors);

    let listener = TcpListener::bind("127.0.0.1:8000").await.unwrap();

    axum::serve(listener, app).await.unwrap();
}
