---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/medium
- monster/type/construct
aliases: ["Clockwork Oaken Bolter"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 80, Mordenkainen's Tome of Foes p. 126
---
# [Clockwork Oaken Bolter](3-Mechanics\CLI\bestiary\construct/clockwork-oaken-bolter-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 80, Mordenkainen's Tome of Foes p. 126*  

```statblock
"name": "Clockwork Oaken Bolter (MPMM)"
"size": "Medium"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "117"
"hit_dice": "18d8 + 36"
"stats":
- !!int "12"
- !!int "18"
- !!int "15"
- !!int "3"
- !!int "10"
- !!int "1"
"speed": "30 ft."
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "understands one language of its creator but can't speak"
"cr": "5"
"traits":
- "desc": "The clockwork has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The clockwork doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The clockwork makes two Lancing Bolt attacks or one Lancing Bolt attack\
    \ and one Harpoon attack."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft.\
    \ or range 100/400 ft., one target. Hit: dice:2d10 + 4|text(15) (2d10 + 4)\
    \ piercing damage."
  "name": "Lancing Bolt"
- "desc": "Ranged Weapon Attack: dice: d20+7 (+7) to hit, range 50/200 ft.,\
    \ one target. Hit: dice:1d10 + 4|text(9) (1d10 + 4) piercing damage, and\
    \ the target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape\
    \ DC 12). While [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) in this\
    \ way, a creature's speed isn't reduced, but it can move only in directions that\
    \ bring it closer to the clockwork. A creature takes dice:1d10|text(5) (1d10)\
    \ slashing damage if it escapes from the grapple or if it tries and fails. The\
    \ clockwork can grapple only one creature at a time with its harpoon."
  "name": "Harpoon"
- "desc": "The clockwork launches an explosive charge at a point within 120 feet.\
    \ Each creature in a 20-foot-radius sphere centered on that point must make a\
    \ DC 15 Dexterity saving throw, taking dice:5d6|text(17) (5d6) fire damage\
    \ on a failed save, or half as much damage on a successful one."
  "name": "Explosive Bolt (Recharge 5-6)"
"bonus_actions":
- "desc": "The clockwork pulls the creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by its Harpoon up to 20 feet closer."
  "name": "Reel In"
"source":
- "MPMM"
- "MTF"
- "RtG"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Clockwork%20Oaken%20Bolter.webp"
```
^statblock

```encounter-table
name: Clockwork Oaken Bolter
creatures:
 - 1: Clockwork Oaken Bolter
```

## Environment

forest, grassland, hill, mountain