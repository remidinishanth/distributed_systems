---
layout: page
title: "Distributed Systems Learning Hub"
---

# Welcome to the Distributed Systems Learning Hub

This repository contains a comprehensive collection of distributed systems concepts, research papers, implementations, and learning resources. Whether you're a student, researcher, or practitioner, you'll find valuable insights into the world of distributed computing.

## 🚀 Quick Start

- **[Browse All Topics]({{ '/topics/' | relative_url }})** - Explore our complete collection of distributed systems topics
- **[Recent Updates](#recent-updates)** - See what's been added recently
- **[Featured Topics](#featured-topics)** - Start with these essential concepts

## 📚 What You'll Find Here

### Core Concepts
- **Consensus Algorithms** (Raft, Paxos)
- **Distributed Storage** (GFS, Cassandra, DynamoDB)
- **Message Processing** (Kafka, MapReduce)
- **Coordination Services** (Zookeeper, etcd, Chubby)

### System Architectures
- **Microservices** and service discovery
- **Caching** strategies and implementations
- **Load Balancing** and traffic management
- **Networking** fundamentals

### Practical Implementations
- **Container Orchestration** (Kubernetes, Docker)
- **Cloud Services** (AWS, monitoring, storage)
- **Development Tools** (Protobuf, build systems)

## 📊 Repository Overview

{% assign all_content_files = site.pages | where_exp: "page", "page.path contains '.md'" | where_exp: "page", "page.path != 'topics.md'" | where_exp: "page", "page.path != 'index.md'" | where_exp: "page", "page.path != 'about.md'" | where_exp: "page", "page.name != 'JEKYLL_SETUP.md'" | where_exp: "page", "page.name != 'PR_INSTRUCTIONS.md'" %}

<div class="repository-stats">
  <div class="stat-card">
    <h3>{{ all_content_files | size }}</h3>
    <p>Total Files</p>
  </div>
  <div class="stat-card">
    {% assign unique_categories = all_content_files | map: "category" | uniq | size %}
    <h3>{{ unique_categories }}</h3>
    <p>Categories</p>
  </div>
  <div class="stat-card">
    {% assign files_by_dir = all_content_files | group_by_exp: "page", "page.path | split: '/' | pop | join: '/'" %}
    <h3>{{ files_by_dir | size }}</h3>
    <p>Directories</p>
  </div>
</div>

## 🎯 Featured Topics

<div class="featured-topics">
{% assign featured = "raft,kafka,kubernetes,databases,google_file_system,blockchain_learnings" | split: "," %}
{% for topic_name in featured %}
  {% assign topic_pages = all_content_files | where_exp: "page", "page.path contains topic_name" %}
  {% for topic in topic_pages limit: 1 %}
  <div class="featured-topic">
    <h3><a href="{{ topic.url | relative_url }}">{{ topic.title | default: topic_name | replace: "_", " " | replace: "-", " " | capitalize }}</a></h3>
    <p>{{ topic.description | default: "Essential distributed systems concept" }}</p>
  </div>
  {% endfor %}
{% endfor %}
</div>

## 📖 Learning Path

### Beginner
1. Start with [Networking Basics]({{ '/networking/basics_5_layer/' | relative_url }})
2. Understand [Distributed Systems Fundamentals](#)
3. Learn about [Consensus Algorithms]({{ '/raft/' | relative_url }})

### Intermediate
1. Explore [Database Systems]({{ '/databases/' | relative_url }})
2. Study [Caching Strategies]({{ '/caching/' | relative_url }})
3. Understand [Message Queues]({{ '/kafka/' | relative_url }})

### Advanced
1. Deep dive into [Kubernetes]({{ '/kubernetes/' | relative_url }})
2. Study [Google File System]({{ '/google_file_system/' | relative_url }})
3. Explore [Blockchain Technologies]({{ '/blockchain_learnings/' | relative_url }})

## 🔗 External Resources

- **MIT 6.824 Distributed Systems** - [Course Notes](https://wizardforcel.gitbooks.io/distributed-systems-engineering-lecture-notes/content/l01-intro.html)
- **High Scalability Blog** - [http://highscalability.com/](http://highscalability.com/)
- **Papers We Love** - Research papers in computer science
- **CNCF Landscape** - [Cloud Native Technologies](https://landscape.cncf.io/)

## 🤝 Contributing

This is a living document that grows with new learnings and discoveries. Feel free to:
- Add new topics and concepts
- Improve existing documentation
- Share interesting papers and resources
- Suggest better organization

---

*Last updated: {{ site.time | date: "%B %d, %Y" }}*

<style>
.featured-topics {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
  margin: 2rem 0;
}

.featured-topic {
  border: 1px solid #e1e4e8;
  border-radius: 6px;
  padding: 1rem;
  background: linear-gradient(135deg, #f6f8fa 0%, #fff 100%);
}

.featured-topic h3 {
  margin-top: 0;
  margin-bottom: 0.5rem;
}

.featured-topic h3 a {
  color: #24292e;
  text-decoration: none;
}

.featured-topic h3 a:hover {
  color: #0366d6;
}

.featured-topic p {
  color: #586069;
  margin-bottom: 0;
  font-size: 0.9em;
}

.repository-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 1rem;
  margin: 2rem 0;
}

.stat-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 1.5rem;
  border-radius: 8px;
  text-align: center;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.stat-card h3 {
  font-size: 2.5rem;
  margin: 0 0 0.5rem 0;
  font-weight: bold;
  color: white;
  border: none;
}

.stat-card p {
  margin: 0;
  font-size: 0.9rem;
  opacity: 0.9;
}
</style>
