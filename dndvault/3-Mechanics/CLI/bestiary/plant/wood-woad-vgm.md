---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/forest
- monster/size/medium
- monster/type/plant
aliases: ["Wood Woad"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 198, Dragon of Icespire Peak, Sleeping Dragon's Wake
---
# [Wood Woad](3-Mechanics\CLI\bestiary\plant/wood-woad-vgm.md)
*Source: Volo's Guide to Monsters p. 198, Dragon of Icespire Peak, Sleeping Dragon's Wake*  

A wood woad is a powerful plant in humanoid form invested with the soul of someone who gave up life to become an eternal guardian.

## Born of Sacrifice

The ritual to create a wood woad is a primeval secret passed down through generations of savage societies and dark druid circles. Performing the ritual isn't necessarily an act of evil, if the victim-to-be has entered into a bargain that requires it to be a willing sacrifice.

In the ritual a living person's chest is pierced and the heart removed. A seed is then pushed into the heart, and it is placed in a tree. Any hollow or crook will do, but often a special cavity is carved out of the trunk. The tree is then bathed and watered with the blood of the sacrificed victim, and the body is buried among the tree's roots. After three days, a sprout emerges from the ground at the base of the tree and swiftly grows into a humanoid form.

This new body, armored in tough bark and bearing a gnarled club and shield, is at once ready to perform its duty. The one who performed the ritual sets the wood woad to its task, and the creature follows those orders unceasingly.

Wood woads are drawn to creatures that have close ties to nature, and that protect and respect the land, such as druids and treants. Some treants have wood woad servants by virtue of age-old pacts with druids or fey that performed the rituals, while others acquire the services of freed wood woads that find renewed purpose in the domain of a kindred guardian.

## Pitiless Protectors

A wood woad has a hole where its heart would be, just as does the body of its former self, buried in the earth. Those who become wood woads trade their free will and all sense of sentiment for supernatural strength and a deathless duty. They exist only to protect woodlands and the people who tend them. A wood woad's face is void and expressionless, except for the motes of light that swim about in its eye sockets. Wood woads speak little, and when not being called upon to take action, they root themselves in the earth and silently take sustenance from it.

## Uprooted by Immortality

Like a tree, a wood woad needs only sunlight, air, and nutrients from the earth to go on living. Because they are undying, some wood woads outlive their original purpose. The site a wood woad guards might lose its power or significance over time, or those whom it was assigned to guard might themselves die. If it is freed from its specific duties, a wood woad might roam to find another place of natural beauty or fey influence to watch over.

```statblock
"name": "Wood Woad (VGM)"
"size": "Medium"
"type": "plant"
"alignment": "Lawful Neutral"
"ac": !!int "18"
"ac_class": "natural armor, [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "75"
"hit_dice": "10d8 + 30"
"stats":
- !!int "18"
- !!int "12"
- !!int "16"
- !!int "10"
- !!int "13"
- !!int "8"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Athletics": !!int "7"
  "Stealth": !!int "4"
  "Perception": !!int "4"
"damage_vulnerabilities": "fire"
"damage_resistances": "bludgeoning, piercing"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "Sylvan"
"cr": "5"
"traits":
- "desc": "In the wood woad's hand, its club is magical and deals dice:3d4|text(7)\
    \ (3d4) extra damage (included in its attacks)."
  "name": "Magic Club"
- "desc": "The wood woad has advantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks it makes in any terrain with ample obscuring plant life."
  "name": "Plant Camouflage"
- "desc": "The wood woad regains 10 hit points at the start of its turn if it is in\
    \ contact with the ground. If the wood woad takes fire damage, this trait doesn't\
    \ function at the start of the wood woad's next turn. The wood woad dies only\
    \ if it starts its turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
- "desc": "Once on each of its turns, the wood woad can use 10 feet of its movement\
    \ to step magically into one living tree within 5 feet of it and emerge from a\
    \ second living tree within 60 feet of it that it can see, appearing in an unoccupied\
    \ space within 5 feet of the second tree. Both trees must be Large or bigger."
  "name": "Tree Stride"
"actions":
- "desc": "The wood woad makes two attacks with its club."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d4 + 4|text(14) (4d4 + 4) bludgeoning damage."
  "name": "Club"
"source":
- "VGM"
- "DIP"
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Wood%20Woad.webp"
```
^statblock

```encounter-table
name: Wood Woad
creatures:
 - 1: Wood Woad
```

## Environment

forest