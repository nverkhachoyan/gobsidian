---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/11
- monster/environment/coastal
- monster/environment/mountain
- monster/environment/underdark
- monster/size/large
- monster/type/aberration
aliases: ["Balhannoth"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 119
---
# [Balhannoth](3-Mechanics\CLI\bestiary\aberration/balhannoth-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 119*  

> [!quote]- A quote from Mordenkainen  
> 
> There are no virtues in the Shadowfell. Thanks to the balhannoth, even hope is punished with death.

## Balhannoth

Native to the Shadowfell, the vicious, predatory balhannoth alters reality in its lair to make the place appear inviting to travelers. Once they step inside, the balhannoth springs its trap.

## False Hope

Thanks to a limited form of telepathy, a balhannoth can sense the desires of other creatures and identify images of places where they expect those desires to be met. The balhannoth then warps reality around it, remaking its environment so that it matches the appearance of the place the creature seeks.

The balhannoth never quite gets all the details right, and plenty of incongruities might give away the deception, but the imitation is good enough to fool desperate creatures into stumbling into the monster's clutches.

## Malevolent Entities

A balhannoth thrives on fear and despair, taking pleasure in the horror its victims experience. It terrorizes its prey by using its reality—warping powers to mask its presence until it can snatch the target. Then it teleports away to feed on its victims.

## Useful Slaves

Drow hunting parties and other denizens of the Underdark sometimes venture into the Shadowfell to capture balhannoths. They install the creatures as guardians, protecting passages from enemy intruders and cutting off avenues of retreat or watching over slaves.

## A Balhannoth's Lair

In the Shadowfell, balhannoths make their lairs near places inhabited by creatures they hunt. They typically haunt well-traveled roads and paths, snatching people who come along. A balhannoth that has been captured and exploited by drow might lair in caves near Underdark passages and guard the ways in and out of a drow enclave.

```statblock
"name": "Balhannoth (MTF)"
"size": "Large"
"type": "aberration"
"alignment": "Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "114"
"hit_dice": "12d10 + 48"
"stats":
- !!int "17"
- !!int "8"
- !!int "18"
- !!int "6"
- !!int "15"
- !!int "8"
"speed": "25 ft., climb 25 ft."
"saves":
  "Constitution": !!int "8"
"skillsaves":
  "Perception": !!int "6"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)"
"senses": "blindsight 500 ft. (blind beyond this radius), passive Perception 16"
"languages": "understands Deep Speech, telepathy 1 mile"
"cr": "11"
"traits":
- "desc": "If the balhannoth fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (2/Day)"
"actions":
- "desc": "The balhannoth makes a bite attack and up to two tentacle attacks, or it\
    \ makes up to four tentacle attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d10 + 3|text(25) (4d10 + 3) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) bludgeoning damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 15)\
    \ and is moved up to 5 feet toward the balhannoth. Until this grapple ends, the\
    \ target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained), and\
    \ the balhannoth can't use this tentacle against other targets. The balhannoth\
    \ has four tentacles."
  "name": "Tentacle"
"legendary_actions":
- "desc": "The balhannoth makes one bite attack against one creature it has [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)."
  "name": "Bite Attack"
- "desc": "The balhannoth magically teleports, along with any equipment it is wearing\
    \ or carrying and any creatures it has [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
    \ up to 60 feet to an unoccupied space it can see."
  "name": "Teleport"
- "desc": "The balhannoth magically becomes [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ for up to 10 minutes or until immediately after it makes an attack roll."
  "name": "Vanish"
"lair_actions":
- "desc": "When fighting inside its lair, a balhannoth can use lair actions. On initiative\
    \ count 20 (losing initiative ties), a balhannoth can take one lair action to\
    \ cause one of the following effects; the balhannoth can't use the same lair action\
    \ two rounds in a row:"
  "name": ""
- "desc": "- The balhannoth warps reality around it in an area up to 500 feet square.\
    \ After 10 minutes, the terrain in the area reshapes to assume the appearance\
    \ of a location sought by one intelligent creature whose mind the balhannoth has\
    \ read (see Regional Effects below). The transformation affects nonliving material\
    \ only and can't create anything with moving parts or magical properties. Any\
    \ object created in this area is, upon close inspection, revealed as a fake. Books\
    \ are filled with empty pages, golden items are obvious counterfeits, and so on.\
    \ The transformation lasts until the balhannoth dies or uses this lair action\
    \ again.  \n- The balhannoth targets one creature within 500 feet of it. The target\
    \ must succeed on a DC 16 Wisdom saving throw or the target, along with whatever\
    \ it is wearing and carrying, teleports to an unoccupied space of the balhannoth's\
    \ choice within 60 feet of it.  \n- The balhannoth targets one creature within\
    \ 500 feet of it. The target must succeed on a DC 16 Wisdom saving throw or the\
    \ balhannoth becomes [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ to that creature for 1 minute. This effect ends if the balhannoth attacks the\
    \ target.  "
  "name": ""
"regional_effects":
- "desc": "A region containing a balhannoth's lair becomes warped by the creature's\
    \ unnatural presence, which creates one or more of the following effects:"
  "name": ""
- "desc": "- Creatures within 1 mile of the balhannoth's lair experience a sensation\
    \ of being close to whatever they desire most. The sensation grows stronger the\
    \ closer the creatures come to the balhannoth's lair.  \n- The balhannoth can\
    \ sense the strongest desires of any humanoid within 1 mile of it and learns whether\
    \ those desires involve a place: a safe location to rest, a temple, home, or somewhere\
    \ else.  "
  "name": ""
- "desc": "If the balhannoth dies, these effects end immediately."
  "name": ""
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Balhannoth.webp"
```
^statblock

```encounter-table
name: Balhannoth
creatures:
 - 1: Balhannoth
```

## Environment

coastal, mountain, underdark