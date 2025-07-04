---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/8
- monster/environment/forest
- monster/environment/swamp
- monster/environment/urban
- monster/size/large
- monster/type/plant
aliases: ["Corpse Flower"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 82, Mordenkainen's Tome of Foes p. 127
---
# [Corpse Flower](3-Mechanics\CLI\bestiary\plant/corpse-flower-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 82, Mordenkainen's Tome of Foes p. 127*  

```statblock
"name": "Corpse Flower (MPMM)"
"size": "Large"
"type": "plant"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "12"
"hp": !!int "127"
"hit_dice": "15d10 + 45"
"stats":
- !!int "14"
- !!int "14"
- !!int "16"
- !!int "7"
- !!int "15"
- !!int "3"
"speed": "20 ft., climb 20 ft."
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 120 ft. (blind beyond this radius), passive Perception 12"
"languages": ""
"cr": "8"
"traits":
- "desc": "When first encountered, a corpse flower contains the corpses of dice:\
    \ 1d6 + 3|avg|noform (1d6 + 3) Humanoids. A corpse flower can hold the remains\
    \ of up to nine Humanoids. These remains have total cover against attacks and\
    \ other effects outside the corpse flower. If the corpse flower dies, the corpses\
    \ within it can be pulled free."
  "name": "Corpses"
- "desc": "The corpse flower can climb difficult surfaces, including upside down on\
    \ ceilings, without needing to make an ability check."
  "name": "Spider Climb"
- "desc": "Each creature that starts its turn within 10 feet of the corpse flower\
    \ or one of its zombies must make a DC 14 Constitution saving throw, unless the\
    \ creature is a Construct or an Undead. On a failed save, the creature is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until the start of its next turn. On a successful save, the creature is immune\
    \ to the Stench of Death of all corpse flowers for 24 hours."
  "name": "Stench of Death"
"actions":
- "desc": "The corpse flower makes three Tentacle attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) bludgeoning damage plus dice:3d6|text(10)\
    \ (3d6) poison damage."
  "name": "Tentacle"
- "desc": "The corpse flower swallows one unsecured Humanoid corpse within 10 feet\
    \ of it, along with any equipment the corpse is wearing or carrying."
  "name": "Harvest the Dead"
"bonus_actions":
- "desc": "The corpse flower digests one corpse in its body and instantly regains\
    \ dice:2d10|text(11) (2d10) hit points. Nothing of the digested corpse remains.\
    \ Any equipment on the corpse is expelled from the corpse flower in its space."
  "name": "Digest"
- "desc": "The corpse flower animates one corpse in its body, turning it into a [zombie](/3-Mechanics/CLI/bestiary/undead/zombie.md).\
    \ The zombie appears in an unoccupied space within 5 feet of the corpse flower\
    \ and acts immediately after it in the initiative order. The zombie acts as an\
    \ ally of the corpse flower but isn't under its control, and the flower's stench\
    \ clings to it (see Stench of Death)."
  "name": "Reanimate"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Corpse%20Flower.webp"
```
^statblock

```encounter-table
name: Corpse Flower
creatures:
 - 1: Corpse Flower
```

## Environment

forest, swamp, urban