#!/usr/bin/env ruby

require 'yaml'

def validate_yaml_frontmatter(file_path)
  content = File.read(file_path)
  
  # Check if file starts with front matter
  unless content.start_with?('---')
    puts "❌ #{file_path}: No front matter found"
    return false
  end
  
  # Extract front matter
  parts = content.split('---', 3)
  if parts.length < 3
    puts "❌ #{file_path}: Incomplete front matter"
    return false
  end
  
  front_matter = parts[1]
  
  begin
    yaml_data = YAML.load(front_matter)
    if yaml_data.nil? || !yaml_data.is_a?(Hash)
      puts "❌ #{file_path}: Invalid YAML structure"
      return false
    end
    
    # Check required fields
    unless yaml_data['layout']
      puts "❌ #{file_path}: Missing 'layout' field"
      return false
    end
    
    unless yaml_data['title']
      puts "❌ #{file_path}: Missing 'title' field"
      return false
    end
    
    puts "✅ #{file_path}: Valid YAML front matter"
    return true
    
  rescue Psych::SyntaxError => e
    puts "❌ #{file_path}: YAML syntax error - #{e.message}"
    return false
  end
end

def scan_markdown_files(directory = '.')
  valid_count = 0
  invalid_count = 0
  
  Dir.glob(File.join(directory, '**', '*.md')).each do |file_path|
    # Skip files that should not have front matter
    next if file_path.include?('/_') || file_path.include?('/.git')
    next if File.basename(file_path).start_with?('_')
    next if ['JEKYLL_SETUP.md', 'PR_INSTRUCTIONS.md'].include?(File.basename(file_path))
    
    if validate_yaml_frontmatter(file_path)
      valid_count += 1
    else
      invalid_count += 1
    end
  end
  
  puts "\n📊 Summary:"
  puts "✅ Valid files: #{valid_count}"
  puts "❌ Invalid files: #{invalid_count}"
  
  return invalid_count == 0
end

# Run the validation
if __FILE__ == $0
  puts "🔍 Validating YAML front matter in markdown files..."
  puts "=" * 50
  
  success = scan_markdown_files
  
  if success
    puts "\n🎉 All markdown files have valid YAML front matter!"
    exit 0
  else
    puts "\n💥 Some files have invalid YAML front matter!"
    exit 1
  end
end
