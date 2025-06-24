use axum::{Router, extract::Path};
use serde::Serialize;

pub fn test1() -> String {
    format!("Hello from utils")
}
