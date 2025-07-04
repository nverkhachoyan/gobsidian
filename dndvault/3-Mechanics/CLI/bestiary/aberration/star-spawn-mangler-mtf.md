---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/5
- monster/size/medium
- monster/type/aberration
aliases: ["Star Spawn Mangler"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 236, Divine Contention, Dragon of Icespire Peak, Storm Lord's Wrath
---
# [Star Spawn Mangler](3-Mechanics\CLI\bestiary\aberration/star-spawn-mangler-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 236, Divine Contention, Dragon of Icespire Peak, Storm Lord's Wrath*  

A mangler is a low-slung, creeping horror with multiple gangly arms. A mangler most often has six arms, but one can have any number from four to eight.

Manglers creep along the ground or the walls, sticking to shadows, hiding in spots that seem too shallow or well-lit to conceal anything. They appear smaller than their true size, thanks to their hunched posture and emaciated frame. Cultists summon these creatures to serve as guards and assassins, two roles at which they excel.

## Star Spawn

> [!quote]- A quote from Mordenkainen  
> 
> Stars don't spawn these creatures. Such beautiful lights shouldn't be blamed for such balefulness.

The Material Plane represents only one small part of the multiverse. Beyond the best-known planes of existence lie realms that are lethal to mortal life. Some are so hostile that even a moment's contact with such a place is enough to plunge a mortal mind into madness. Yet beings do exist that are native to these realms: beings that are eternally hungering, searching, warring, sometimes dreaming. These Elder Evils are far older than most of the mortal races and always horrific to humanoid minds.

However much they might desire to enter and dominate the Material Plane, the Elder Evils are unable or unwilling to leave their realms. Some are imprisoned in their dimensions by external forces, some are inextricably bound to their home realities, and others simply can't find any way out.

### Heralds of Doom

The creatures known as star spawn are the heralds, servants, foot soldiers, and lieutenants of the Elder Evils, capable of taking on forms that can journey to the Material Plane. They arrive most often in the wake of a comet—or perhaps such a phenomenon merely signals that star spawn are in the vicinity and available for communication. When the signs are right, warlocks and cultists hasten to gather together, read aloud their blasphemous texts, and conduct the mind-searing rituals that guide the blazing star spawn into the world.

> [!quote]- A quote from Mordenkainen  
> 
> The cultists who blaspheme reality by calling out to Elder Evils often speak of a Far Realm from which these entities hail. In truth, there is no one place or space from which they come. There is the multiverse of things that are, and there is the multiverse of things that shouldn't be.

### Elder Evil Blessings

Disciples of certain Elder Evils can bestow supernatural gifts on those who serve that cult, including star spawn. The following powers are unique to specific cults; typically a creature has only one.

- Cult of Atropus, the World Born Dead  
- Cult of Borem, of the Lake of Boiling Mud  
- Cult of Haask, the Voice of Hargut  
- Cult of Tharizdun, the Chained God  
- Cult of Tyranthraxus, the Flamed One  

```statblock
"name": "Star Spawn Mangler (MTF)"
"size": "Medium"
"type": "aberration"
"alignment": "Chaotic Evil"
"ac": !!int "14"
"hp": !!int "71"
"hit_dice": "13d8 + 13"
"stats":
- !!int "8"
- !!int "18"
- !!int "12"
- !!int "11"
- !!int "12"
- !!int "7"
"speed": "40 ft., climb 40 ft."
"saves":
  "Dexterity": !!int "7"
  "Constitution": !!int "4"
"skillsaves":
  "Stealth": !!int "7"
"damage_resistances": "cold"
"damage_immunities": "psychic"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Deep Speech"
"cr": "5"
"traits":
- "desc": "On the first round of each combat, the mangler has advantage on attack\
    \ rolls against a creature that hasn't taken a turn yet."
  "name": "Ambusher"
- "desc": "While in dim light or darkness, the mangler can take the Hide action as\
    \ a bonus action."
  "name": "Shadow Stealth"
"actions":
- "desc": "The mangler makes two claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage. If the attack roll\
    \ has advantage, the target also takes dice:2d6|text(7) (2d6) psychic damage."
  "name": "Claw"
- "desc": "The mangler makes six claw attacks against one target. Either before or\
    \ after these attacks, it can move up to its speed as a bonus action without provoking\
    \ opportunity attacks."
  "name": "Flurry of Claws (Recharge 4-6)"
"source":
- "MTF"
- "DC"
- "DIP"
- "SLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Star%20Spawn%20Mangler.webp"
```
^statblock

```encounter-table
name: Star Spawn Mangler
creatures:
 - 1: Star Spawn Mangler
```