---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
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
BookSource: Mordenkainen's Tome of Foes p. 127
---
# [Corpse Flower](3-Mechanics\CLI\bestiary\plant/corpse-flower-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 127*  

> [!quote]- A quote from Mordenkainen  
> 
> Corpse-flower seedlings are quite useful for various purposes. Simply kill and bury a necromancer, and you should have a good crop in about a week.

## Corpse Flower

A corpse flower can sprout atop the grave of an evil necromancer or the remains of powerful undead. Unless it is uprooted and burned while it is still a seedling, the corpse flower grows to enormous size over several weeks, then tears itself free of the earth and begins scavenging humanoid corpses from battlefields and graveyards. Using its fibrous tentacles, it stuffs the remains into its body and feeds on carrion to repair itself. The plant has a malevolent bent and despises the living.

## Horrible Odor

With or without humanoid corpses nested in its body, a corpse flower exudes a stench of decay that can overwhelm the senses of nearby creatures, causing them to become nauseated. The stench, which serves as a defense mechanism, fades `dice: 2d4|avg|noform` (`2d4`) days after the corpse flower dies.

```statblock
"name": "Corpse Flower (MTF)"
"size": "Large"
"type": "plant"
"alignment": "Chaotic Evil"
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
  \ [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened)"
"senses": "blindsight 120 ft. (blind beyond this radius), passive Perception 12"
"languages": ""
"cr": "8"
"traits":
- "desc": "When first encountered, a corpse flower contains the corpses of dice:\
    \ 1d6 + 3|avg|noform (1d6 + 3) humanoids. A corpse flower can hold the remains\
    \ of up to nine dead humanoids. These remains have total cover against attacks\
    \ and other effects outside the corpse flower. If the corpse flower dies, the\
    \ corpses within it can be pulled free.\n\nWhile it has at least one humanoid\
    \ corpse in its body, the corpse flower can use a bonus action to do one of the\
    \ following:\n\n- The corpse flower digests one humanoid corpse in its body and\
    \ instantly regains dice:2d10|text(11) (2d10) hit points. Nothing of the digested\
    \ body remains. Any equipment on the corpse is expelled from the corpse flower\
    \ in its space.  \n- The corpse flower animates one dead humanoid in its body,\
    \ turning it into a zombie. The zombie appears in an unoccupied space within 5\
    \ feet of the corpse flower and acts immediately after it in the initiative order.\
    \ The zombie acts as an ally of the corpse flower but isn't under its control,\
    \ and the flower's stench clings to it (see the Stench of Death trait).  "
  "name": "Corpses"
- "desc": "The corpse flower can climb difficult surfaces, including upside down on\
    \ ceilings, without needing to make an ability check."
  "name": "Spider Climb"
- "desc": "Each creature that starts its turn within 10 feet of the corpse flower\
    \ or one of its zombies must make a DC 14 Constitution saving throw, unless the\
    \ creature is a construct or undead. On a failed save, the creature is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ until the end of the turn. Creatures that are immune to poison damage or the\
    \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) condition automatically\
    \ succeed on this saving throw. On a successful save, the creature is immune to\
    \ the stench of all corpse flowers for 24 hours."
  "name": "Stench of Death"
"actions":
- "desc": "The corpse flower makes three tentacle attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) bludgeoning damage, and the target\
    \ must succeed on a DC 14 Constitution saving throw or take dice:4d6|text(14)\
    \ (4d6) poison damage."
  "name": "Tentacle"
- "desc": "The corpse flower grabs one unsecured dead humanoid within 10 feet of it\
    \ and stuffs the corpse into itself, along with any equipment the corpse is wearing\
    \ or carrying. The remains can be used with the Corpses trait."
  "name": "Harvest the Dead"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Corpse%20Flower.webp"
```
^statblock

```encounter-table
name: Corpse Flower
creatures:
 - 1: Corpse Flower
```

## Environment

forest, swamp, urban