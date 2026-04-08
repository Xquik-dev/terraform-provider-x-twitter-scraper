resource "x-twitter-scraper_compose" "example_compose" {
  step = "compose"
  additional_context = "https://x.com/elonmusk/status/1234567890"
  call_to_action = "Follow for more"
  draft = "AI is changing everything. Here\'s why."
  goal = "engagement"
  has_link = false
  has_media = false
  media_type = "none"
  style_username = "elonmusk"
  tone = "professional"
  topic = "AI trends in 2025"
}
