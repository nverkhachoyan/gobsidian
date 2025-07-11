function fileTreeComponent() {
  return {
    expanded: false,
    fileTreeData: window.fileTreeData,
    expandedNodes: new Set(),
    
    init() {
      console.log('File tree component initialized', this.fileTreeData);
      // Store reference to this component globally for click handlers
      window.fileTreeInstance = this;
    },
    
    toggleNode(path) {
      if (this.expandedNodes.has(path)) {
        this.expandedNodes.delete(path);
      } else {
        this.expandedNodes.add(path);
      }
      // Force re-render by updating the data
      this.fileTreeData = { ...this.fileTreeData };
    },
    
    isExpanded(path) {
      return this.expandedNodes.has(path);
    },
    
    renderTree(node) {
      if (!node) return '';
      
      let html = '<ul class="file-tree__list">';
      
      // Render folders (children)
      if (node.children && node.children.length > 0) {
        node.children.forEach(child => {
          const isExpanded = this.isExpanded(child.path);
          html += `
            <li class="file-tree__item file-tree__item--folder ${isExpanded ? 'expanded' : ''}">
              <div class="file-tree__folder-header">
                <span onclick="window.fileTreeInstance.toggleNode('${child.path}')" class="file-tree__toggle">
                  <i class="fas fa-folder file-tree__icon file-tree__icon--closed"></i>
                  <i class="fas fa-folder-open file-tree__icon file-tree__icon--open"></i>
                </span>
                
                <span onclick="window.fileTreeInstance.toggleNode('${child.path}')" class="file-tree__file_name file-tree__toggle">
                  ${child.name}
                </span>

                <a href="/${child.path}/" class="file-tree__link file-tree__folder-icon" title="Go to folder">
                  <i class="fa-solid fa-arrow-up-right-from-square"></i>
                </a>
              </div>
              <div class="file-tree__children">
                ${isExpanded ? this.renderTree(child) : ''}
              </div>
            </li>
          `;
        });
      }
      
      // Render files
      if (node.files && node.files.length > 0) {
        node.files.forEach(file => {
          html += `
            <li class="file-tree__item file-tree__item--file">
              <i class="far fa-file-alt file-tree__icon"></i>
              <a href="${file.url}" class="file-tree__link">${file.name}</a>
            </li>
          `;
        });
      }
      
      html += '</ul>';
      return html;
    }
  }
}
