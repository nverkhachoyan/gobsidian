---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/2
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/large
- monster/type/beast
aliases: ["Aurochs"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 207
---
# [Aurochs](3-Mechanics\CLI\bestiary\beast/aurochs-vgm.md)
*Source: Volo's Guide to Monsters p. 207*  

Bahgtru, son of Gruumsh and Luthic, is the orc deity of unbridled strength. Legend says Bahgtru needed a mount as fierce as him for making war, so he sought a mighty aurochs, subjugated the creature with his bare hands, and hauled it to Nishrek, Gruumsh's realm. Bahgtru named the beast Kazaht, or "Bull" in Orc. On Kazaht's bare back, Bahgtru charges into battle, ramming into an enemy host and leaping over the aurochs's horns to land in the midst of his foes.

Orcs that revere Bahgtru might tend a stable of war bulls that carry them into combat. Trained to be fierce mounts from a young age, aurochs are sacred symbols of Bahgtru. No orc will eat such creatures, which are treated as honored warriors when they perish.

```statblock
"name": "Aurochs (VGM)"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "11"
"ac_class": "natural armor"
"hp": !!int "38"
"hit_dice": "4d10 + 16"
"stats":
- !!int "20"
- !!int "10"
- !!int "19"
- !!int "2"
- !!int "12"
- !!int "5"
"speed": "50 ft."
"senses": "passive Perception 11"
"languages": ""
"cr": "2"
"traits":
- "desc": "If the aurochs moves at least 20 feet straight toward a target and then\
    \ hits it with a gore attack on the same turn, the target takes an extra dice:2d8|text(9)\
    \ (2d8) piercing damage. If the target is a creature, it must succeed on a DC\
    \ 15 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Charge"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 5|text(14) (2d8 + 5) piercing damage."
  "name": "Gore"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Aurochs.webp"
```
^statblock

```encounter-table
name: Aurochs
creatures:
 - 1: Aurochs
```

## Environment

mountain, grassland, hill