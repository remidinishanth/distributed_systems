#!/usr/bin/env ruby

# Script to add front matter to markdown files that don't have it
# This helps Jekyll process all markdown files properly

require 'fileutils'

def add_frontmatter_to_file(file_path)
  content = File.read(file_path)
  
  # Skip if file already has front matter
  return if content.start_with?('---')
  
  # Extract title from filename or first heading
  filename = File.basename(file_path, '.md')
  title = filename.gsub(/[-_]/, ' ').split.map(&:capitalize).join(' ')
  
  # Try to extract title from first heading in content
  if content.match(/^#\s+(.+)$/m)
    title = $1.strip
  end
  
  # Determine category from folder structure
  folder_path = File.dirname(file_path)
  category = folder_path == '.' ? 'general' : File.basename(folder_path)
  
  # Generate front matter
  front_matter = <<~FRONTMATTER
    ---
    layout: page
    title: "#{title}"
    category: "#{category}"
    ---

  FRONTMATTER
  
  # Write updated content
  File.write(file_path, front_matter + content)
  puts "Added front matter to: #{file_path}"
end

def process_directory(dir_path = '.')
  Dir.glob(File.join(dir_path, '**', '*.md')).each do |file_path|
    # Skip files that start with underscore or are in special directories
    next if file_path.include?('/_') || file_path.include?('/.git')
    next if File.basename(file_path).start_with?('_')
    next if ['index.md', 'topics.md'].include?(File.basename(file_path))
    
    add_frontmatter_to_file(file_path)
  end
end

# Run the script
if __FILE__ == $0
  puts "Adding front matter to markdown files..."
  process_directory
  puts "Done!"
end
