---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/6
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/large
- monster/type/monstrosity
aliases: ["Chimera"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 39, Sleeping Dragon's Wake. Available in the SRD and the Basic Rules.
---
# [Chimera](3-Mechanics\CLI\bestiary\monstrosity/chimera.md)
*Source: Monster Manual p. 39, Sleeping Dragon's Wake. Available in the SRD and the Basic Rules.*  

Chimeras were created after mortals summoned Demogorgon to the world. The Prince of Demons, unimpressed with the creatures that surrounded it, transformed them into horrific, multi-headed monstrosities. This act gave rise to the first chimeras.

Gifted with demonic cruelty, a chimera serves as a grim reminder of what happens when demon princes find their way to the Material Plane. A typical specimen has the hindquarters of a large goat, the forequarters of a lion, and the leathery wings of a dragon, along with the heads of all three of those creatures. The monster likes to surprise its victims, swooping down from the sky and engulfing prey with its fiery breath before landing.

## Conflicted Creature

A chimera combines the worst aspects of its three parts. Its dragon head drives it to raid, plunder, and accumulate a great hoard. Its leonine nature compels it to hunt and kill powerful creatures that threaten its territory. Its goat head grants it a vicious, stubborn streak that compels it to fight to the death. These three aspects drive a chimera to stake out a territory that is as large as 10 miles wide. It preys on wild game, viewing more powerful creatures as rivals to be humiliated and defeated. Its greatest rivals are dragons, griffons, manticores, perytons, and wyverns. When it hunts, the chimera looks for easy ways to amuse itself. It enjoys the fear and suffering of weaker creatures. The monster often toys with its prey, breaking off an attack prematurely and leaving a creature wounded and terrified before returning to finish it off.

## Servant of Evil

Though chimeras are far from cunning, their draconic ego makes them susceptible to flattery and gifts. If offered food and treasure, a chimera might spare a traveler. A villain can lure a chimera into service by keeping it well fed and its treasure hoard well stocked.

```statblock
"name": "Chimera"
"size": "Large"
"type": "monstrosity"
"alignment": "Chaotic Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "114"
"hit_dice": "12d10 + 48"
"stats":
- !!int "19"
- !!int "11"
- !!int "19"
- !!int "3"
- !!int "14"
- !!int "10"
"speed": "30 ft., fly 60 ft."
"skillsaves":
  "Perception": !!int "8"
"senses": "darkvision 60 ft., passive Perception 18"
"languages": "understands Draconic but can't speak"
"cr": "6"
"actions":
- "desc": "The chimera makes three attacks: one with its bite, one with its horns,\
    \ and one with its claws. When its fire breath is available, it can use the breath\
    \ in place of its bite or horns."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d12 + 4|text(10) (1d12 + 4) bludgeoning damage."
  "name": "Horns"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) slashing damage."
  "name": "Claws"
- "desc": "The dragon head exhales fire in a 15-foot cone. Each creature in that area\
    \ must make a DC 15 Dexterity saving throw, taking dice:7d8|text(31) (7d8)\
    \ fire damage on a failed save, or half as much damage on a successful one."
  "name": "Fire Breath (Recharge 5-6)"
"source":
- "MM"
- "PotA"
- "RoT"
- "SKT"
- "TftYP"
- "WDMM"
- "SDW"
- "BGDIA"
- "DoDk"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Chimera.webp"
```
^statblock

```encounter-table
name: Chimera
creatures:
 - 1: Chimera
```

## Environment

underdark, mountain, grassland, hill