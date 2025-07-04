---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-2
- monster/size/small
- monster/type/elemental
aliases: ["Magmin"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Monster Manual p. 212. Available in the SRD.
---
# [Magmin](3-Mechanics\CLI\bestiary\elemental/magmin.md)
*Source: Monster Manual p. 212. Available in the SRD.*  

A grinning, mischievous magmin resembles a stumpy humanoid sculpted from a black shell of magma. Even when it isn't ablaze and radiating heat like a bonfire, small jets of flame erupt from its porous skin.

## Summoned Pyromaniacs

Magmins are fire elemental spirits bound into physical forms by magic, and they appear in the Material Plane only when summoned. They view flammable objects as kindling for a grand conflagration, and only the magical control exerted by their summoners keeps them from setting everything they touch ablaze. Their propensity for fire and havoc makes them ideal for spreading chaos and destruction. A mob of magmins summoned inside a castle can reduce it to a burning shell within minutes.

### Fiery Destruction

Although its flame is potent, the magmin's hard magma shell prevents it from instantly igniting everything it comes into contact with. However, like the fires inside them, magmins are capricious and unpredictable. Moreover, as simple elemental creations, they are oblivious to the harm their native element causes creatures of the Material Plane.

If it has the opportunity while in service to its master, a magmin seeks out areas of great heat, such as forest fires or the bubbling magma of an active volcano. At other times, a magmin compulsively looses fire from its fingertips, delighting in setting objects ablaze.

```statblock
"name": "Magmin"
"size": "Small"
"type": "elemental"
"alignment": "Chaotic Neutral"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "9"
"hit_dice": "2d6 + 2"
"stats":
- !!int "7"
- !!int "15"
- !!int "12"
- !!int "8"
- !!int "11"
- !!int "10"
"speed": "30 ft."
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "fire"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Ignan"
"cr": "1/2"
"traits":
- "desc": "When the magmin dies, it explodes in a burst of fire and magma. Each creature\
    \ within 10 feet of it must make a DC 11 Dexterity saving throw, taking dice:2d6|text(7)\
    \ (2d6) fire damage on a failed save, or half as much damage on a successful\
    \ one. Flammable objects that aren't being worn or carried in that area are ignited."
  "name": "Death Burst"
- "desc": "As a bonus action, the magmin can set itself ablaze or extinguish its flames.\
    \ While ablaze, the magmin sheds bright light in a 10-foot radius and dim light\
    \ for an additional 10 feet."
  "name": "Ignited Illumination"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6|text(7) (2d6) fire damage. If the target is a creature or\
    \ a flammable object, it ignites. Until a creature takes an action to douse the\
    \ fire, the target takes dice:1d6|text(3) (1d6) fire damage at the end of\
    \ each of its turns."
  "name": "Touch"
"source":
- "MM"
- "PotA"
- "SKT"
- "ToA"
- "BGDIA"
- "WBtW"
- "SatO"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Magmin.webp"
```
^statblock

```encounter-table
name: Magmin
creatures:
 - 1: Magmin
```