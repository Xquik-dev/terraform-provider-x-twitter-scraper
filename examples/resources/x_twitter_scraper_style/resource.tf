# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

resource "x-twitter-scraper_style" "example_style" {
  id    = "id"
  label = "Professional Voice"
  tweets = [{
    text = "Excited to share our latest research findings."
  }]
}
