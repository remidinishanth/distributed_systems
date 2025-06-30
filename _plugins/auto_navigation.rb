# Auto Navigation Generator Plugin
# This plugin automatically generates navigation based on folder structure

Jekyll::Hooks.register :site, :post_read do |site|
  # Generate navigation structure from pages
  navigation = {}
  
  site.pages.each do |page|
    next if page.path.start_with?('_') || page.path == 'index.md' || page.path == 'topics.md'
    next unless page.path.end_with?('.md')
    
    # Extract folder structure
    path_parts = page.path.split('/')
    current_level = navigation
    
    # Build nested structure
    path_parts[0..-2].each do |folder|
      current_level[folder] ||= { 'pages' => [], 'subfolders' => {} }
      current_level = current_level[folder]['subfolders']
    end
    
    # Add page to appropriate folder
    folder_name = path_parts.length > 1 ? path_parts[-2] : 'root'
    if path_parts.length > 1
      parent_folder = navigation
      path_parts[0..-2].each do |folder|
        parent_folder = parent_folder[folder]['subfolders'] if parent_folder[folder]
      end
      if parent_folder[folder_name]
        parent_folder[folder_name]['pages'] << {
          'title' => page.data['title'] || page.basename.gsub(/[-_]/, ' ').split.map(&:capitalize).join(' '),
          'url' => page.url,
          'description' => page.data['description']
        }
      end
    else
      navigation['root'] ||= { 'pages' => [], 'subfolders' => {} }
      navigation['root']['pages'] << {
        'title' => page.data['title'] || page.basename.gsub(/[-_]/, ' ').split.map(&:capitalize).join(' '),
        'url' => page.url,
        'description' => page.data['description']
      }
    end
  end
  
  # Store in site data for use in templates
  site.data['auto_navigation'] = navigation
end
