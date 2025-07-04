---
obsidianUIMode: preview
cssclasses:
  - json5e-item
tags:
  - item
aliases: 
SourceType: Magic Item
NoteIcon: magicitem
BookSource: 
weight: 0
cost: "0"
---

<% await tp.file.move("/3-Mechanics/Items/" + tp.file.title) %>

<%*
const hasTitle = !tp.file.title.startsWith("NewMagicItem");
let title;
if (!hasTitle) {
    title = await tp.system.prompt("Enter Item Name");
    await tp.file.rename(title);
} else {
    title = tp.file.title;
}
_%>


# `=this.file.name`
*Melee Weapon, uncommon (requires attunement by a druid or ranger)*  

- **Damage**: 1d4 S
- **Properties**: [Light](/2-Mechanics/CLI/rules/item-properties.md#Light), Requires Attunement
- **Cost**: `=this.cost`
- **Weight**: `=this.weight`

Description

*Source: `=this.BookSource`*