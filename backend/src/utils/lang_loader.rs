use libloading::Library;
use std::collections::HashMap;
use std::path::Path;

use crate::utils::lang_cacher::LANG_CACHE;

fn load_language(name: &str, lang: &str) -> Result<Language, Box<dyn std::error::Error>> {}
