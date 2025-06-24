use once_cell::sync::Lazy;
use std::collections::HashMap;
use std::sync::Mutex;
use tree_sitter::Language;

pub static LANG_CACHE: Lazy<Mutex<HashMap<String, Language>>> =
    Lazy::new(|| Mutex::new(HashMap::new()));
