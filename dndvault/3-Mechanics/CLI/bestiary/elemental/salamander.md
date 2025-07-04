---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/5
- monster/environment/underdark
- monster/size/large
- monster/type/elemental
aliases: ["Salamander"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Monster Manual p. 266. Available in the SRD.
---
# [Salamander](3-Mechanics\CLI\bestiary\elemental/salamander.md)
*Source: Monster Manual p. 266. Available in the SRD.*  

Salamanders slither across the Sea of Ash on the Elemental Plane of Fire, their sinuous coils and jagged spines smoldering. Intense heat washes off their bodies, while their yellow eyes glow like candles in the deep-set hollows of their hawkish faces.

Salamanders adore power, and they delight in setting fire to things. Outside their home plane, they play among the burning skeletons of charred trees as forest fires rage around them, or slither down the slopes of erupting volcanoes to linger in fire pits and magma floes.

## Fire Snakes

Salamanders hatch from eggs that are two-foot-diameter spheres of smoldering obsidian. When a salamander is ready to hatch, it melts its way through the egg's thick shell and emerges as a fire snake. A fire snake matures into a salamander adult within a year.

## Slaves of the Efreet

Long ago, the efreet hired azers to build the fabled City of Brass, but then failed in their attempt to enslave that mystical race when the azers' work was done. Turning instead to strike against the salamanders, the efreet had better luck in establishing a slave race, which they use to unleash war and destruction across the planes.

Salamanders despise the azers, believing that if the efreet had succeeded in dominating that race of elemental crafters, the salamanders would still be free. The efreet use this enmity to their own advantage, stoking the salamanders' hatred and pitting them against the efreets' former servants.

The efreet suffer salamanders to serve no other master; when efreet encounter salamanders dedicated to the cults of Elemental Evil, they slay them rather than taking them as slaves.

## Domineering Nobles

Although salamanders follow the destructive impulses of their fiery nature, slavery under the efreet has impacted the culture of free salamanders. They rule their own societies according to the efreet model, in which larger and stronger salamanders claim dominion over their lesser kin.

As salamanders age, they increase in size and status, rising to positions of power as cruel nobles among their kind. Nobles rule wandering bands of salamanders, which move across the Elemental Plane of Fire like desert nomads, raiding other communities for treasure.

## Living Forges

Salamanders generate intense heat, and when they fight, their weapons glow red and sear the bodies of their enemies on contact. Even approaching a salamander is dangerous, since flesh blisters and burns in its proximity.

This inherent heat is an asset to salamanders' skill as smiths, allowing them to soften and shape iron and steel with their bare hands. Although not as meticulous as azers, salamanders number among the greatest metalsmiths in all the planes. Powerful creatures summon them as warriors, but others enlist the salamanders for their crafting skills, or bind them to forges and ovens to generate limitless heat.

```statblock
"name": "Salamander"
"size": "Large"
"type": "elemental"
"alignment": "Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "90"
"hit_dice": "12d10 + 24"
"stats":
- !!int "18"
- !!int "14"
- !!int "15"
- !!int "11"
- !!int "10"
- !!int "12"
"speed": "30 ft."
"damage_vulnerabilities": "cold"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "fire"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Ignan"
"cr": "5"
"traits":
- "desc": "A creature that touches the salamander or hits it with a melee attack while\
    \ within 5 feet of it takes dice:2d6|text(7) (2d6) fire damage."
  "name": "Heated Body"
- "desc": "Any metal melee weapon the salamander wields deals an extra dice:1d6|text(3)\
    \ (1d6) fire damage on a hit (included in the attack)."
  "name": "Heated Weapons"
"actions":
- "desc": "The salamander makes two attacks: one with its spear and one with its tail."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:2d6 + 4|text(11) (2d6 + 4) piercing\
    \ damage, or dice:2d8 + 4|text(13) (2d8 + 4) piercing damage if used with\
    \ two hands to make a melee attack, plus dice:1d6|text(3) (1d6) fire damage."
  "name": "Spear"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) bludgeoning damage plus dice:2d6|text(7)\
    \ (2d6) fire damage, and the target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 14). Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ the salamander can automatically hit the target with its tail, and the salamander\
    \ can't make tail attacks against other targets."
  "name": "Tail"
"source":
- "MM"
- "PotA"
- "SKT"
- "TftYP"
- "ToA"
- "BGDIA"
- "JttRC"
- "DoSI"
- "KftGV"
- "GotSF"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Salamander.webp"
```
^statblock

```encounter-table
name: Salamander
creatures:
 - 1: Salamander
```

## Environment

underdark