---
NoteIcon: region 
Tags: Category/Region
aliases:
 - some alias
commonraces:
 - Human
alignments:
 - all
commonraces:
 - Humans
 - Elves
 - Dwarves
region:
 - Some area
largestcity:
 - Big City
languages:
 - Common
religions:
 - Lathander
---

<% await tp.file.move("/2-World/" + tp.file.title) %>

<%*
const hasTitle = !tp.file.title.startsWith("NewRegion");
let title;
if (!hasTitle) {
    title = await tp.system.prompt("Location Name");
    await tp.file.rename(title);
} else {
    title = tp.file.title;
}
_%>


> [!infobox]
> # `=this.file.name`
> ![[MapPlaceholder.png|cover hsmall]]
> ###### Geography
> Type |  Stat |
> ---|---|
> Aliases|`=this.aliases`|
> Region | `=this.region` |
> Largest City|`=this.largestcity`|
> ###### Society 
> Type |  Stat |
> ---|---|
> Races|`=this.commonraces`|
> Languages|`=this.languages`|
> Religions|`=this.religions`|
>Alignments|`=this.alignments`|

# `=this.file.name`
## Overview
Placeholder

### Placeholder Map
![[MapPlaceholder.png|Placeholder Map]]

### Placeholder Picture
![[ImagePlaceholder.png|Placeholder Picture]]

Placeholder

## Geography
Placeholder

## Government
Placeholder

## Society
Placeholder

## Trade
Placeholder

##  Defenses
Placeholder

## History
Placeholder

## Rumors and Legends
Placeholder

## Inhabitants
Placeholder

## Additional Details
Placeholder

