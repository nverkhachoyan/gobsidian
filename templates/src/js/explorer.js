export default function explorerComponent() {
  return {
    expanded: false,
    fileTreeData: null,
    expandedNodes: new Set(),
    
    init() {
      Alpine.effect(() => {
        this.fileTreeData = Alpine.store('fileTree').data;
      })
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
      
      let html = '<ul class="explorer__list">';
      
      // Render folders (children)
      if (node.children && node.children.length > 0) {
        node.children.forEach(child => {
          const isExpanded = this.isExpanded(child.path);
          html += `
            <li class="explorer__item explorer__item--folder ${isExpanded ? 'expanded' : ''}">
              <div class="explorer__folder-header">
                <span onclick="window.fileTreeInstance.toggleNode('${child.path}')" class="explorer__toggle">
                  <i class="fas fa-folder explorer__icon explorer__icon--closed"></i>
                  <i class="fas fa-folder-open explorer__icon explorer__icon--open"></i>
                </span>
                
                <span onclick="window.fileTreeInstance.toggleNode('${child.path}')" class="explorer__file_name explorer__toggle">
                  ${child.name}
                </span>

                <a href="/${child.path}/" class="explorer__link explorer__folder-icon" title="Go to folder">
                  <i class="fa-solid fa-arrow-up-right-from-square"></i>
                </a>
              </div>
              <div class="explorer__children">
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
            <li class="explorer__item explorer__item--file">
              <i class="far fa-file-alt explorer__icon"></i>
              <a href="${file.url}" class="explorer__link">${file.name}</a>
            </li>
          `;
        });
      }
      
      html += '</ul>';
      return html;
    }
  }
}
