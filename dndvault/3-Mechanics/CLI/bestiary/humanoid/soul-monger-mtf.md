---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/11
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Soul Monger"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 226
---
# [Soul Monger](3-Mechanics\CLI\bestiary\humanoid/soul-monger-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 226*  

## Soul Monger

Wracked with despair over the loss of memories of a brighter time, soul mongers now crave the vitality of others. The aching void within a soul monger radiates outward, manifesting as an unbearable weight that drains the vigor of anyone unfortunate enough to be in its presence. Those who have escaped the onslaught of a soul monger can hardly shake the memory of the sound it makes—the moan of a tortured soul, lost in a bottomless well of tragedy.

## Shadar-kai

In the perpetual gloom of the Shadowfell lives a society that serves the Raven Queen. They were brought into that dusky realm in ages past, so long ago that they're now perfectly adapted to that cheerless environment, both physically and mentally.

## Soul Custodians

Shadar-kai watch over both the Shadowfell and the material world, scouting out choice souls and tragedies that might please their deity. They are rumored to be able to coax worldly events along tragic paths for her amusement. The Raven Queen is famously cryptic even to her most devoted followers, however. The shadar-kai's efforts are rewarded only with vague omens they interpret as best they can.

## Blighted Elves

Shadar-kai were once elves, but eons of exposure to the debilitating influence of the Shadowfell has left them joyless and mournful. In that realm, they have the appearance of withered elves: pale hair, wrinkled gray skin, and swollen joints give them a corpselike aspect. They appear more youthful while on other planes, but their skin always retains its deathly pallor. They dress in dark cloaks and heavy veils, detest mirrors, and avoid keeping things that remind them of their age.

```statblock
"name": "Soul Monger (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Neutral"
"ac": !!int "15"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "123"
"hit_dice": "19d8 + 38"
"stats":
- !!int "8"
- !!int "17"
- !!int "14"
- !!int "19"
- !!int "15"
- !!int "13"
"speed": "30 ft."
"saves":
  "Charisma": !!int "5"
  "Dexterity": !!int "7"
  "Wisdom": !!int "6"
"skillsaves":
  "Perception": !!int "6"
"damage_immunities": "necrotic, psychic"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 17"
"languages": "Common, Elvish"
"cr": "11"
"traits":
- "desc": "The soul monger's innate spellcasting ability is Intelligence (spell save\
    \ DC 16, dice: d20+8 (+8) to hit with spell attacks). It can innately cast\
    \ the following spells, requiring no material components:\n\nAt will: [chill\
    \ touch](/3-Mechanics/CLI/spells/chill-touch.md), [poison spray](/3-Mechanics/CLI/spells/poison-spray.md)\n\
    \n1/day each: [bestow curse](/3-Mechanics/CLI/spells/bestow-curse.md), [chain\
    \ lightning](/3-Mechanics/CLI/spells/chain-lightning.md), [finger of death](/3-Mechanics/CLI/spells/finger-of-death.md),\
    \ [gaseous form](/3-Mechanics/CLI/spells/gaseous-form.md), [phantasmal killer](/3-Mechanics/CLI/spells/phantasmal-killer.md),\
    \ [seeming](/3-Mechanics/CLI/spells/seeming.md)"
  "name": "Innate Spellcasting"
- "desc": "The soul monger has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put it to sleep."
  "name": "Fey Ancestry"
- "desc": "The soul monger has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "When the soul monger reduces a creature to 0 hit points, the soul monger\
    \ can gain temporary hit points equal to half the creature's hit point maximum.\
    \ While the soul monger has temporary hit points from this ability, it has advantage\
    \ on attack rolls."
  "name": "Soul Thirst"
- "desc": "Any beast or humanoid, other than a shadar-kai, that starts its turn within\
    \ 5 feet of the soul monger has its speed reduced by 20 feet until the start of\
    \ that creature's next turn."
  "name": "Weight of Ages"
"actions":
- "desc": "The soul monger makes two phantasmal dagger attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d4 + 3|text(13) (4d4 + 3) piercing damage plus dice:3d12|text(19)\
    \ (3d12) necrotic damage, and the target has disadvantage on saving throws until\
    \ the start of the soul monger's next turn."
  "name": "Phantasmal Dagger"
- "desc": "The soul monger emits weariness in a 60-foot cube. Each creature in that\
    \ area must make a DC 16 Constitution saving throw. On a failed save, a creature\
    \ takes dice:10d8|text(45) (10d8) psychic damage and suffers 1 level of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion).\
    \ On a successful save, it takes dice:5d8|text(22) (5d8) psychic damage."
  "name": "Wave of Weariness (Recharge 4-6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Soul%20Monger.webp"
```
^statblock

```encounter-table
name: Soul Monger
creatures:
 - 1: Soul Monger
```

## Environment

underdark, urban