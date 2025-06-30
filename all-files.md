---
layout: page
title: "All Files"
permalink: /all-files/
description: "Complete listing of every markdown file in the repository"
---

# 📁 Complete File Listing

This page shows **every single markdown file** in the repository, organized and searchable.

{% include search_all_files.html %}

{% assign all_content_files = site.pages | where_exp: "page", "page.path contains '.md'" | where_exp: "page", "page.path != 'topics.md'" | where_exp: "page", "page.path != 'index.md'" | where_exp: "page", "path != 'about.md'" | where_exp: "page", "path != 'all-files.md'" | where_exp: "page", "name != 'JEKYLL_SETUP.md'" | where_exp: "page", "name != 'PR_INSTRUCTIONS.md'" %}

## 📊 File Statistics

<div class="file-overview">
  <div class="overview-stat">
    <span class="stat-number">{{ all_content_files | size }}</span>
    <span class="stat-label">Total Files</span>
  </div>
  
  {% assign files_by_category = all_content_files | group_by: "category" %}
  <div class="overview-stat">
    <span class="stat-number">{{ files_by_category | size }}</span>
    <span class="stat-label">Categories</span>
  </div>
  
  {% assign files_by_dir = all_content_files | group_by_exp: "page", "page.path | split: '/' | pop | join: '/'" %}
  <div class="overview-stat">
    <span class="stat-number">{{ files_by_dir | size }}</span>
    <span class="stat-label">Directories</span>
  </div>
  
  {% assign files_with_descriptions = all_content_files | where_exp: "page", "page.description" %}
  <div class="overview-stat">
    <span class="stat-number">{{ files_with_descriptions | size }}</span>
    <span class="stat-label">With Descriptions</span>
  </div>
</div>

## 🗂️ Files by Category

{% for category_group in files_by_category %}
<div class="category-section">
  <h3 class="category-title">
    {% if category_group.name == "" or category_group.name == nil %}
      📄 Uncategorized
    {% else %}
      📁 {{ category_group.name | capitalize }}
    {% endif %}
    <span class="file-count">({{ category_group.items | size }} files)</span>
  </h3>
  
  <div class="files-grid">
    {% assign sorted_files = category_group.items | sort: "title" %}
    {% for file in sorted_files %}
    <div class="file-card">
      <div class="file-header">
        <h4><a href="{{ file.url | relative_url }}">{{ file.title | default: file.name | replace: ".md", "" | replace: "_", " " | replace: "-", " " | capitalize }}</a></h4>
        <span class="file-path">{{ file.path }}</span>
      </div>
      
      {% if file.description %}
      <p class="file-description">{{ file.description }}</p>
      {% endif %}
      
      <div class="file-meta">
        {% if file.category %}
        <span class="file-category">{{ file.category }}</span>
        {% endif %}
        
        {% if file.tags %}
        <div class="file-tags">
          {% for tag in file.tags %}
          <span class="file-tag">{{ tag }}</span>
          {% endfor %}
        </div>
        {% endif %}
      </div>
    </div>
    {% endfor %}
  </div>
</div>
{% endfor %}

## 📂 Files by Directory Structure

{% for dir_group in files_by_dir %}
<div class="directory-section">
  <h3 class="directory-title">
    {% if dir_group.name == "" %}
      📄 Root Directory
    {% else %}
      📁 {{ dir_group.name }}
    {% endif %}
    <span class="file-count">({{ dir_group.items | size }} files)</span>
  </h3>
  
  <ul class="directory-files">
    {% assign sorted_files = dir_group.items | sort: "name" %}
    {% for file in sorted_files %}
    <li class="directory-file">
      <a href="{{ file.url | relative_url }}" class="file-link">
        <span class="file-icon">📝</span>
        <span class="file-name">{{ file.name | replace: ".md", "" }}</span>
        {% if file.title and file.title != file.name %}
        <span class="file-title">{{ file.title }}</span>
        {% endif %}
      </a>
      {% if file.description %}
      <p class="file-desc">{{ file.description }}</p>
      {% endif %}
    </li>
    {% endfor %}
  </ul>
</div>
{% endfor %}

<style>
.file-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 1rem;
  margin: 2rem 0;
  padding: 1.5rem;
  background: #f6f8fa;
  border-radius: 8px;
  border: 1px solid #e1e4e8;
}

.overview-stat {
  text-align: center;
  padding: 1rem;
  background: white;
  border-radius: 6px;
  border: 1px solid #e1e4e8;
}

.stat-number {
  display: block;
  font-size: 2rem;
  font-weight: bold;
  color: #0366d6;
}

.stat-label {
  display: block;
  font-size: 0.9rem;
  color: #586069;
  margin-top: 0.5rem;
}

.category-section, .directory-section {
  margin: 2rem 0;
  border: 1px solid #e1e4e8;
  border-radius: 8px;
  overflow: hidden;
}

.category-title, .directory-title {
  background: #f6f8fa;
  margin: 0;
  padding: 1rem;
  border-bottom: 1px solid #e1e4e8;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.file-count {
  font-size: 0.9rem;
  color: #586069;
  font-weight: normal;
}

.files-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1rem;
  padding: 1rem;
}

.file-card {
  border: 1px solid #e1e4e8;
  border-radius: 6px;
  padding: 1rem;
  background: white;
  transition: box-shadow 0.2s;
}

.file-card:hover {
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.file-header h4 {
  margin: 0 0 0.5rem 0;
}

.file-header h4 a {
  color: #0366d6;
  text-decoration: none;
}

.file-header h4 a:hover {
  text-decoration: underline;
}

.file-path {
  font-size: 0.8rem;
  color: #586069;
  font-family: monospace;
}

.file-description {
  color: #24292e;
  font-size: 0.9rem;
  line-height: 1.4;
  margin: 0.5rem 0;
}

.file-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-top: 0.5rem;
}

.file-category {
  background: #f1f8ff;
  color: #0366d6;
  padding: 0.2rem 0.5rem;
  border-radius: 3px;
  font-size: 0.8rem;
}

.file-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.25rem;
}

.file-tag {
  background: #e1f5fe;
  color: #0277bd;
  padding: 0.1rem 0.4rem;
  border-radius: 2px;
  font-size: 0.7rem;
}

.directory-files {
  list-style: none;
  margin: 0;
  padding: 0;
}

.directory-file {
  border-bottom: 1px solid #f1f3f4;
  padding: 0.75rem 1rem;
}

.directory-file:last-child {
  border-bottom: none;
}

.file-link {
  text-decoration: none;
  color: #24292e;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.file-link:hover {
  color: #0366d6;
}

.file-icon {
  font-size: 1rem;
}

.file-name {
  font-weight: 500;
}

.file-title {
  color: #586069;
  font-weight: normal;
}

.file-desc {
  margin: 0.5rem 0 0 2rem;
  color: #586069;
  font-size: 0.9rem;
  line-height: 1.4;
}

@media (max-width: 768px) {
  .files-grid {
    grid-template-columns: 1fr;
  }
  
  .file-overview {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .file-link {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .file-desc {
    margin-left: 0;
  }
}
</style>
