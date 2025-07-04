---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/1-4
- monster/environment/mountain
- monster/environment/swamp
- monster/size/small
- monster/type/aberration
aliases: ["Star Spawn Grue"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 234
---
# [Star Spawn Grue](3-Mechanics\CLI\bestiary\aberration/star-spawn-grue-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 234*  

Fanged and lipless, the ever-grinning, madly staring grue lopes about on spindly legs and long arms. Bristles and spines project from odd patches of its pallid skin, and it's long fingers end in broken and dirty nails. Grues are the weakest of the star spawn. A host of writhing, scrambling grues typically accompanies more powerful star spawn. Their constant chittering and shrieking produces discordant psychic energy that disrupts thought patterns in other creatures. Intelligent creatures experience flashing colors, hallucinations, disorientation, and waves of hopelessness when they find themselves near a group of star spawn grues.

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
"name": "Star Spawn Grue (MTF)"
"size": "Small"
"type": "aberration"
"alignment": "Lawful Evil"
"ac": !!int "11"
"hp": !!int "17"
"hit_dice": "5d6"
"stats":
- !!int "6"
- !!int "13"
- !!int "10"
- !!int "9"
- !!int "11"
- !!int "6"
"speed": "30 ft."
"damage_immunities": "psychic"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Deep Speech"
"cr": "1/4"
"traits":
- "desc": "Creatures within 20 feet of the grue that aren't aberrations have disadvantage\
    \ on saving throws, as well as on attack rolls against creatures other than a\
    \ star spawn grue."
  "name": "Aura of Madness"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 1|text(6) (2d4 + 1) piercing damage, and the target must\
    \ succeed on a DC 10 Wisdom saving throw or attack rolls against it have advantage\
    \ until the start of the grue's next turn."
  "name": "Confounding Bite"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Star%20Spawn%20Grue.webp"
```
^statblock

```encounter-table
name: Star Spawn Grue
creatures:
 - 1: Star Spawn Grue
```

## Environment

mountain, swamp