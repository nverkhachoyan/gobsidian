---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/5
- monster/environment/desert
- monster/environment/mountain
- monster/size/large
- monster/type/elemental
aliases: ["Air Elemental"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Monster Manual p. 124, Dragon of Icespire Peak, Storm Lord's Wrath, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.
---
# [Air Elemental](3-Mechanics\CLI\bestiary\elemental/air-elemental.md)
*Source: Monster Manual p. 124, Dragon of Icespire Peak, Storm Lord's Wrath, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.*  

An air elemental is a funneling cloud of whirling air with a vague semblance of a face. Although it likes to race across the ground, picking up dust and debris as it goes, it can also fly and attack from above.

An air elemental can turn itself into a screaming cyclone, creating a whirlwind that batters creatures even as it flings them away.

## Elementals

Elementals are incarnations of the elements that make up the universe: air, earth, fire, and water. Though little more than animated energy on their own planes of existence, they can be called on by spellcasters and powerful beings to take shape and perform tasks.

### Living Elements

On its home plane, an elemental is a bodiless life force. Its dim consciousness manifests as a physical shape only when focused by the power of magic. A wild spirit of elemental force has no desire except to course through the element of its native plane. Like beasts of the Material Plane, these elemental spirits have no society or culture, and little sense of being.

### Conjured by Magic

Certain spells and magic items can conjure an elemental, summoning it from the Inner Planes to the Material Plane. Elementals instinctively resent being pulled from their native planes and bound into service. A creature that summons an elemental must assert force of will to control it.

### Bound and Shaped

Powerful magic can bind an elemental spirit into a material template that defines a specific use and function. Invisible stalkers are air elementals bound to a specific form, in the same way that water elementals can be shaped into water weirds.

### Elemental Nature

An elemental doesn't require air, food, drink, or sleep.

```statblock
"name": "Air Elemental"
"size": "Large"
"type": "elemental"
"alignment": "Neutral"
"ac": !!int "15"
"hp": !!int "90"
"hit_dice": "12d10 + 24"
"stats":
- !!int "14"
- !!int "20"
- !!int "14"
- !!int "6"
- !!int "10"
- !!int "6"
"speed": "fly 90 ft. (hover)"
"damage_resistances": "lightning; thunder; bludgeoning, piercing, slashing from nonmagical\
  \ attacks"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
  \ [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Auran"
"cr": "5"
"traits":
- "desc": "The elemental can enter a hostile creature's space and stop there. It can\
    \ move through a space as narrow as 1 inch wide without squeezing."
  "name": "Air Form"
"actions":
- "desc": "The elemental makes two slam attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 5|text(14) (2d8 + 5) bludgeoning damage."
  "name": "Slam"
- "desc": "Each creature in the elemental's space must make a DC 13 Strength saving\
    \ throw. On a failure, a target takes dice:3d8 + 2|text(15) (3d8 + 2) bludgeoning\
    \ damage and is flung up 20 feet away from the elemental in a random direction\
    \ and knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone). If a thrown\
    \ target strikes an object, such as a wall or floor, the target takes dice:1d6|text(3)\
    \ (1d6) bludgeoning damage for every 10 feet it was thrown. If the target is\
    \ thrown at another creature, that creature must succeed on a DC 13 Dexterity\
    \ saving throw or take the same damage and be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\n\
    \nIf the saving throw is successful, the target takes half the bludgeoning damage\
    \ and isn't flung away or knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Whirlwind (Recharge 4-6)"
"source":
- "MM"
- "HotDQ"
- "PotA"
- "RoT"
- "SKT"
- "TftYP"
- "ToA"
- "WDMM"
- "GoS"
- "DIP"
- "SLW"
- "ERLW"
- "IMR"
- "MOT"
- "TCE"
- "CM"
- "WBtW"
- "JttRC"
- "DSotDQ"
- "KftGV"
- "GotSF"
- "BMT"
- "DoDk"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Air%20Elemental.webp"
```
^statblock

```encounter-table
name: Air Elemental
creatures:
 - 1: Air Elemental
```

## Environment

mountain, desert