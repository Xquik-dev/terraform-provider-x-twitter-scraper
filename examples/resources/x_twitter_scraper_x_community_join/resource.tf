resource "x-twitter-scraper_x_community_join" "example_x_community_join" {
  account         = "@example"
  idempotency_key = "terraform-join-community-v1"
  target_id       = "1500000000000000000"
}
