# Jekyll GitHub Pages Setup

This repository is now configured to automatically build and deploy to GitHub Pages using Jekyll with recursive building of all markdown files.

## 🚀 What's Been Set Up

### 1. Jekyll Configuration (`_config.yml`)
- Site metadata and description
- Plugin configuration for SEO, feeds, and sitemaps
- Custom collections and defaults
- Markdown processing with syntax highlighting

### 2. GitHub Actions Workflow (`.github/workflows/jekyll-gh-pages.yml`)
- Automatic building on push to main branch
- Custom Ruby/Jekyll setup with plugin support
- Deployment to GitHub Pages

### 3. Site Structure
```
├── _config.yml              # Jekyll configuration
├── _layouts/                 # Custom page layouts
│   ├── default.html         # Main site layout with navigation
│   └── topic.html           # Layout for topic pages
├── _includes/               # Reusable components
│   ├── head.html           # HTML head with meta tags
│   ├── navigation.html     # Site navigation
│   └── footer.html         # Site footer
├── _plugins/               # Custom Jekyll plugins
│   └── auto_navigation.rb  # Auto-generates navigation from folders
├── _scripts/               # Utility scripts
│   └── add_frontmatter.rb  # Adds YAML front matter to markdown files
├── assets/                 # CSS and other assets
│   └── custom.css          # Custom styling
├── index.md                # Homepage
├── topics.md               # Auto-generated topics index
├── about.md                # About page
└── Gemfile                 # Ruby dependencies
```

### 4. Automatic Features
- **Recursive Building**: All markdown files in subdirectories are automatically processed
- **Auto Navigation**: Topics are automatically organized by folder structure
- **Front Matter**: YAML metadata added to all markdown files for proper processing
- **Responsive Design**: Mobile-friendly layout
- **Syntax Highlighting**: Code blocks with proper highlighting
- **Math Support**: MathJax for mathematical expressions
- **SEO Optimization**: Meta tags, sitemaps, and structured data

## 🔧 Configuration Steps

### 1. Update Repository Settings
1. Go to your GitHub repository settings
2. Navigate to "Pages" section
3. Set source to "GitHub Actions"
4. The site will be available at `https://[username].github.io/[repository-name]`

### 2. Customize Site Information
Edit `_config.yml` to update:
```yaml
title: "Your Site Title"
description: "Your site description"
baseurl: "/your-repository-name"  # Only if not using username.github.io
url: "https://your-username.github.io"
```

### 3. Add Front Matter to Existing Files
The script `_scripts/add_frontmatter.rb` has already been run to add YAML front matter to your markdown files. Each file now has:
```yaml
---
layout: page
title: "Page Title"
category: "folder-name"
tags: ["tag1", "tag2"]
description: "Page description"
---
```

## 📝 Adding New Content

### For New Topics
1. Create a new markdown file in an appropriate folder
2. Add front matter at the top:
```yaml
---
layout: page
title: "Your Topic Title"
category: "topic-category"
tags: ["tag1", "tag2"]
description: "Brief description of the topic"
---
```
3. Write your content in markdown
4. The page will automatically appear in the topics index

### For New Categories
1. Create a new folder
2. Add a `readme.md` file in the folder with front matter
3. The category will automatically appear in navigation

## 🎨 Customization

### Styling
- Edit `assets/custom.css` for visual customizations
- Modify `_layouts/` files for structural changes
- Update `_includes/navigation.html` for menu changes

### Navigation
- The navigation is automatically generated from your folder structure
- Manual navigation items can be added in `_includes/navigation.html`

### Homepage
- Edit `index.md` to customize the homepage content
- Featured topics are automatically pulled from specific folders

## 🚀 Local Development

To run the site locally:

```bash
# Install dependencies
bundle install

# Serve the site locally
bundle exec jekyll serve

# Site will be available at http://localhost:4000
```

## 📊 Features Included

- ✅ Automatic site generation from markdown files
- ✅ Responsive design for mobile and desktop
- ✅ Search functionality
- ✅ Syntax highlighting for code blocks
- ✅ Math equation support (MathJax)
- ✅ SEO optimization
- ✅ RSS feed generation
- ✅ Sitemap generation
- ✅ Social media meta tags
- ✅ Print-friendly styles

## 🔄 Automatic Updates

Every time you push to the main branch:
1. GitHub Actions builds the site
2. All markdown files are processed
3. Navigation is updated automatically
4. Site is deployed to GitHub Pages

Your distributed systems learning hub is now live and will automatically update as you add new content!

## 📚 Next Steps

1. **Verify the site is working**: Check your GitHub Pages URL
2. **Add more content**: Create new markdown files in topic folders
3. **Customize styling**: Modify CSS to match your preferences
4. **Add images**: Store images in topic folders or create an `assets/images/` directory
5. **Enable analytics**: Add Google Analytics tracking if desired

The site will automatically rebuild and deploy whenever you push changes to the main branch.
