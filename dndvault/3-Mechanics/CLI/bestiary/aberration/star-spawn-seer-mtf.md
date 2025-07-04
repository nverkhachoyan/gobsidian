---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/13
- monster/environment/mountain
- monster/environment/swamp
- monster/environment/urban
- monster/size/medium
- monster/type/aberration
aliases: ["Star Spawn Seer"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 236
---
# [Star Spawn Seer](3-Mechanics\CLI\bestiary\aberration/star-spawn-seer-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 236*  

A star spawn seer is most often encountered as the leader of a cult that worships one or more of the Elder Evils. Usually, the seer is the only cult member that grasps the full extent of the horror the cult is venerating.

An entity that appears as a star spawn seer in the Material Plane usually arrives as something different—something disembodied. When a warlock or other spellcaster establishes communication with it, the seer-entity takes control of the mortal's form and spirit, transforming it into a star spawn seer. Whoever the seer once was largely vanishes beneath the corpulent bulk of tumorous skin than builds up in strange whorls all over the seer's body. Hands become bulky, flipper-like appendages capable of grasping their strange staffs—formed of some blend of flesh, bone, and star stuff—but clumsy and painful when used to manipulate other things.

A star spawn seer is almost always accompanied by one or more star spawn hulks. Although the hulk is a worthy combatant in its own right, it's also a vital part of a tactic often used by seers. When a seer deals psychic damage to a hulk, the hulk isn't hurt, while the effect ricochets off the hulk and expands to assault other creatures.

The seer's goal is to tap the energy sources and master the rites that will enable it to extend a bridge between the vulnerable sanity of the Material Plane and the squirming madness of an Elder Evil's prison.

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
"name": "Star Spawn Seer (MTF)"
"size": "Medium"
"type": "aberration"
"alignment": "Lawful Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "153"
"hit_dice": "18d8 + 72"
"stats":
- !!int "14"
- !!int "12"
- !!int "18"
- !!int "22"
- !!int "19"
- !!int "16"
"speed": "30 ft."
"saves":
  "Charisma": !!int "8"
  "Dexterity": !!int "6"
  "Wisdom": !!int "9"
  "Intelligence": !!int "11"
"skillsaves":
  "Perception": !!int "9"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "psychic"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 19"
"languages": "Common, Deep Speech, Undercommon"
"cr": "13"
"traits":
- "desc": "The seer can move through other creatures and objects as if they were difficult\
    \ terrain. Each creature it moves through takes dice:1d10|text(5) (1d10) psychic\
    \ damage; no creature can take this damage more than once per turn. The seer takes\
    \ dice:1d10|text(5) (1d10) force damage if it ends its turn inside an object."
  "name": "Out-of-Phase Movement"
"actions":
- "desc": "The seer makes two comet staff attacks or uses Psychic Orb twice."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 5 ft., one\
    \ target. Hit: dice:1d6 + 6|text(9) (1d6 + 6) bludgeoning damage plus dice:4d8|text(18)\
    \ (4d8) psychic damage, or dice:1d8 + 6|text(10) (1d8 + 6) bludgeoning damage\
    \ plus dice:4d8|text(18) (4d8) psychic damage, if used with two hands, and\
    \ the target must succeed on a DC 19 Constitution saving throw or be [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ until the end of its next turn."
  "name": "Comet Staff"
- "desc": "Ranged Spell Attack: dice: d20+11 (+11) to hit, range 120 feet, one\
    \ target. Hit: dice:5d10|text(27) (5d10) psychic damage."
  "name": "Psychic Orb"
- "desc": "The seer warps space around a creature it can see within 30 feet of it.\
    \ That creature must make a DC 19 Wisdom saving throw. On a failed save, the target,\
    \ along with any equipment it is wearing or carrying, is magically teleported\
    \ up to 60 feet to an unoccupied space the seer can see, and all other creatures\
    \ within 10 feet of the target's original space each takes dice:6d12|text(39)\
    \ (6d12) psychic damage. On a successful save, the target takes dice:3d12|text(19)\
    \ (3d12) psychic damage."
  "name": "Collapse Distance (Recharge 6)"
"reactions":
- "desc": "When the seer would be hit by an attack, it teleports, exchanging positions\
    \ with another star spawn it can see within 60 feet of it. The other star spawn\
    \ is hit by the attack instead."
  "name": "Bend Space"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Star%20Spawn%20Seer.webp"
```
^statblock

```encounter-table
name: Star Spawn Seer
creatures:
 - 1: Star Spawn Seer
```

## Environment

mountain, swamp, urban