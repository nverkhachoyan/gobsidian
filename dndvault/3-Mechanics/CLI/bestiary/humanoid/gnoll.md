---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-2
- monster/environment/desert
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/humanoid/gnoll
aliases: ["Gnoll"]
NoteIcon: monster
BestiaryType: humanoid (gnoll)
SourceType: Bestiary
BookSource: Monster Manual p. 163. Available in the SRD and the Basic Rules.
---
# [Gnoll](3-Mechanics\CLI\bestiary\humanoid/gnoll.md)
*Source: Monster Manual p. 163. Available in the SRD and the Basic Rules.*  

Gnolls are feral humanoids that attack settlements along the frontiers and borderlands of civilization without warning, slaughtering their victims and devouring their flesh.

## Demonic Origin

The origin of the gnolls traces back to a time when the demon lord Yeenoghu found his way to the Material Plane and ran amok. Packs of ordinary hyenas followed in his wake, scavenging the demon lord's kills. Those hyenas were transformed into the first gnolls, parading after Yeenoghu until he was banished back to the Abyss. The gnolls then scattered across the face of the world, a dire reminder of demonic power.

## Nomadic Destroyers

Gnolls are dangerous because they strike at random. They emerge from the wilderness, plunder and slaughter, then move elsewhere. They attack like a plague of locusts, pillaging settlements and leaving little behind but razed buildings, gnawed corpses, and befouled land. Gnolls choose easy targets for their raids. Armored warriors holed up in a fortified castle will survive a rampaging gnoll horde unscathed, even as the towns, villages, and farms that surround the castle are ablaze, their people slaughtered and devoured.

Gnolls rarely build permanent structures or craft anything of lasting value. They don't make weapons or armor, but scavenge such items from the corpses of their fallen victims, stringing ears, teeth, scalps, and other trophies from their foes onto their patchwork armor.

### Thirst for Blood

No goodness or compassion resides in the heart of a gnoll. Like a demon, it lacks anything resembling a conscience, and can't be taught or coerced to put aside its destructive tendencies. The gnolls' frenzied bloodlust makes them an enemy to all, and when they lack a common foe, they fight among themselves. Even the most savage orcs avoid allying with gnolls.

```statblock
"name": "Gnoll"
"size": "Medium"
"type": "humanoid"
"subtype": "gnoll"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "[hide armor](/3-Mechanics/CLI/items/hide-armor.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "22"
"hit_dice": "5d8"
"stats":
- !!int "14"
- !!int "12"
- !!int "11"
- !!int "6"
- !!int "10"
- !!int "7"
"speed": "30 ft."
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Gnoll"
"cr": "1/2"
"traits":
- "desc": "When the gnoll reduces a creature to 0 hit points with a melee attack on\
    \ its turn, the gnoll can take a bonus action to move up to half its speed and\
    \ make a bite attack."
  "name": "Rampage"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing damage."
  "name": "Bite"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing\
    \ damage, or dice:1d8 + 2|text(6) (1d8 + 2) piercing damage if used with two\
    \ hands to make a melee attack."
  "name": "Spear"
- "desc": "Ranged Weapon Attack: dice: d20+3 (+3) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 1|text(5) (1d8 + 1) piercing damage."
  "name": "Longbow"
"source":
- "MM"
- "PotA"
- "TftYP"
- "GoS"
- "BGDIA"
- "ERLW"
- "EGW"
- "IDRotF"
- "SatO"
- "ToFW"
- "BMT"
- "DoDk"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Gnoll.webp"
```
^statblock

```encounter-table
name: Gnoll
creatures:
 - 1: Gnoll
```

## Environment

grassland, forest, hill, desert