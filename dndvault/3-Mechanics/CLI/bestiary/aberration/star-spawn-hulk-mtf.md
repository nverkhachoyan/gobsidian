---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/10
- monster/size/large
- monster/type/aberration
aliases: ["Star Spawn Hulk"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 234
---
# [Star Spawn Hulk](3-Mechanics\CLI\bestiary\aberration/star-spawn-hulk-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 234*  

The hulk is the largest of the known star spawn. Though ogre-like in stature, the hulk's glistening translucent skin reveals a muscled form devoid of an ogre's fat. Pale and seemingly lidless eyes glare balefully from a face distorted by too many teeth and too little nose.

Hulks are seldom encountered without a commanding seer nearby. A hulk appears to have little will of its own, other than to protect its master.

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
"name": "Star Spawn Hulk (MTF)"
"size": "Large"
"type": "aberration"
"alignment": "Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "136"
"hit_dice": "13d10 + 65"
"stats":
- !!int "20"
- !!int "8"
- !!int "21"
- !!int "7"
- !!int "12"
- !!int "9"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "3"
  "Wisdom": !!int "5"
"skillsaves":
  "Perception": !!int "5"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "Deep Speech"
"cr": "10"
"traits":
- "desc": "If the hulk takes psychic damage, each creature within 10 feet of the hulk\
    \ takes that damage instead; the hulk takes none of the damage. In addition, the\
    \ hulk's thoughts and location can't be discerned by magic."
  "name": "Psychic Mirror"
"actions":
- "desc": "The hulk makes two slam attacks. If both attacks hit the same target, the\
    \ target also takes dice:2d8|text(9) (2d8) psychic damage and must succeed\
    \ on a DC 17 Constitution saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the end of the target's next turn."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d8 + 5|text(14) (2d8 + 5) bludgeoning damage."
  "name": "Slam"
- "desc": "The hulk makes a separate slam attack against each creature within 10 feet\
    \ of it. Each creature that is hit must also succeed on a DC 17 Dexterity saving\
    \ throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Reaping Arms (Recharge 5-6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Star%20Spawn%20Hulk.webp"
```
^statblock

```encounter-table
name: Star Spawn Hulk
creatures:
 - 1: Star Spawn Hulk
```