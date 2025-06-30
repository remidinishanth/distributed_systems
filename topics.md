---
layout: page
title: "All Topics"
permalink: /topics/
---

# Distributed Systems Topics

This is a comprehensive collection of distributed systems concepts, implementations, and learning resources. Each topic contains detailed notes, papers, and practical examples.

## Topics by Category

{% assign topics_by_category = site.pages | where_exp: "page", "page.path contains '.md'" | where_exp: "page", "page.path != 'README.md'" | where_exp: "page", "page.path != 'topics.md'" | where_exp: "page", "page.path != 'index.md'" | group_by: "category" %}

{% for category_group in topics_by_category %}
  {% if category_group.name != "" %}
### {{ category_group.name | capitalize }}
  {% else %}
### General Topics
  {% endif %}

  {% assign sorted_topics = category_group.items | sort: "title" %}
  {% for topic in sorted_topics %}
    {% unless topic.path contains '_' or topic.path contains '.github' %}
- [{{ topic.title | default: topic.name | replace: "_", " " | replace: "-", " " | capitalize }}]({{ topic.url | relative_url }})
      {% if topic.description %}
  <br><small>{{ topic.description }}</small>
      {% endif %}
    {% endunless %}
  {% endfor %}
{% endfor %}

## All Topics (Alphabetical)

{% assign all_topics = site.pages | where_exp: "page", "page.path contains '.md'" | where_exp: "page", "page.path != 'README.md'" | where_exp: "page", "page.path != 'topics.md'" | where_exp: "page", "page.path != 'index.md'" | sort: "title" %}

<div class="topics-grid">
{% for topic in all_topics %}
  {% unless topic.path contains '_' or topic.path contains '.github' %}
  <div class="topic-card">
    <h3><a href="{{ topic.url | relative_url }}">{{ topic.title | default: topic.name | replace: "_", " " | replace: "-", " " | capitalize }}</a></h3>
    {% if topic.description %}
      <p>{{ topic.description }}</p>
    {% endif %}
    {% if topic.tags %}
      <div class="tags">
        {% for tag in topic.tags %}
          <span class="tag">{{ tag }}</span>
        {% endfor %}
      </div>
    {% endif %}
  </div>
  {% endunless %}
{% endfor %}
</div>

<style>
.topics-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1.5rem;
  margin-top: 2rem;
}

.topic-card {
  border: 1px solid #e1e4e8;
  border-radius: 6px;
  padding: 1rem;
  background: #fff;
  transition: box-shadow 0.2s;
}

.topic-card:hover {
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.topic-card h3 {
  margin-top: 0;
  margin-bottom: 0.5rem;
}

.topic-card h3 a {
  color: #24292e;
  text-decoration: none;
}

.topic-card h3 a:hover {
  color: #0366d6;
}

.topic-card p {
  color: #586069;
  font-size: 0.9em;
  margin-bottom: 0.5rem;
}

.tags {
  margin-top: 0.5rem;
}

.tag {
  background: #f1f8ff;
  color: #0366d6;
  padding: 0.2rem 0.5rem;
  border-radius: 3px;
  margin-right: 0.5rem;
  font-size: 0.8em;
}
</style>
