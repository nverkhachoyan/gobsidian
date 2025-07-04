---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/21
- monster/size/gargantuan
- monster/type/monstrosity/titan
aliases: ["Astral Dreadnought"]
NoteIcon: monster
BestiaryType: monstrosity (titan)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 117
---
# [Astral Dreadnought](3-Mechanics\CLI\bestiary\monstrosity/astral-dreadnought-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 117*  

> [!quote]- A quote from Mordenkainen  
> 
> Astral dreadnoughts exist for one reason: hubris. Not the hubris of mortals, but the hubris of gods who deem themselves too mighty to be approached and looked upon.

## Astral Dreadnought

Enormous and terrifying monstrosities known as astral dreadnoughts haunt the silvery void of the Astral Plane, causing planar travelers to shudder at the very thought of them. They have been gliding through the astral mists since the dawn of the multiverse, trying to devour all other creatures they encounter.

As big as an [ancient red dragon](/3-Mechanics/CLI/bestiary/dragon/ancient-red-dragon.md) and covered from head to tail in layers of thick, spiked plates, a dreadnought has two gnarled limbs that end in razor-sharp pincer claws. Constellations appear to swirl in the depths of its single eye, and its serpentine, armored tail trails off into the silvery void.

An astral dreadnought lives a solitary existence. On the rare occasion when two dreadnoughts meet, they typically fight until one tires of the conflict and departs. Some mighty villains have enslaved astral dreadnoughts and used them to terrifying effect.

## Antimagic Eye

Astral sailors claim that insanity awaits anyone who gazes into the eye of an astral dreadnought. What one sees reflected in that starry void is the sudden, terrifying realization of one's own mortality. Spellcasters have cause to fear the eye more than others, since it emits a continuous [antimagic field](/3-Mechanics/CLI/spells/antimagic-field.md). The dreadnought can shut off the effect by simply closing its eye, though it seldom has reason to do so.

## Astral Predator

A remorseless, indiscriminate hunter, an astral dreadnought employs terrifying, if unimaginative, tactics. It uses its teeth and claws to tear apart its prey. Instinctively aware of how dangerous spellcasters can be, it maneuvers to keep as many opponents as possible within its antimagic gaze.

An astral dreadnought doesn't have a gullet or a digestive system. Anything it swallows is deposited in a unique demiplane—an enclosed space that contains eons worth of detritus, as well as the remains of dead planar travelers. The place has gravity and breathable air, and organic matter decays there. Although escape from the demiplane is possible with the aid of magic, most creatures arrive here only after they have died. When the dreadnought dies, its demiplane vanishes, and its contents are released into the Astral Plane.

An astral dreadnought doesn't communicate. It simply consumes any prey it finds, then continues its silent patrol. It can't leave the Astral Plane, nor would it want to.

## Titans of the Chained God

Tharizdun, the Chained God, created astral dreadnoughts to devour planar travelers who were seeking portals that lead from the Astral Plane to the Outer Planes—portals they might use to gaze upon their gods or realize some dream of godhood.

Astral dreadnoughts don't procreate, so their population can't grow.

Even though githyanki and other astral voyagers hunt the creatures, they rarely see any success, and the dreadnoughts aren't in danger of becoming extinct anytime soon.

## Titanic Nature

Although it eats and sleeps if it so desires, an astral dreadnought doesn't require air, food, drink, or sleep.

```statblock
"name": "Astral Dreadnought (MTF)"
"size": "Gargantuan"
"type": "monstrosity"
"subtype": "titan"
"alignment": "Unaligned"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "297"
"hit_dice": "17d20 + 119"
"stats":
- !!int "28"
- !!int "7"
- !!int "25"
- !!int "5"
- !!int "14"
- !!int "18"
"speed": "15 ft., fly 80 ft. (hover)"
"saves":
  "Dexterity": !!int "5"
  "Wisdom": !!int "9"
"skillsaves":
  "Perception": !!int "9"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "darkvision 120 ft., passive Perception 19"
"languages": ""
"cr": "21"
"traits":
- "desc": "The astral dreadnought's opened eye creates an area of antimagic, as in\
    \ the [antimagic field](/3-Mechanics/CLI/spells/antimagic-field.md) spell, in\
    \ a 150-foot cone. At the start of each of its turns, the dreadnought decides\
    \ which way the cone faces. The cone doesn't function while the dreadnought's\
    \ eye is closed or while the dreadnought is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)."
  "name": "Antimagic Cone"
- "desc": "The astral dreadnought can't leave the Astral Plane, nor can it be banished\
    \ or otherwise transported out of the Astral Plane."
  "name": "Astral Entity"
- "desc": "Any creature or object that the astral dreadnought swallows is transported\
    \ to a demiplane that can be entered by no other means except a [wish](/3-Mechanics/CLI/spells/wish.md)\
    \ spell or this creature's Donjon Visit ability. A creature can leave the demiplane\
    \ only by using magic that enables planar travel, such as the [plane shift](/3-Mechanics/CLI/spells/plane-shift.md)\
    \ spell. The demiplane resembles a stone cave roughly 1,000 feet in diameter with\
    \ a ceiling 100 feet high. Like a stomach, it contains the remains of the dreadnought's\
    \ past meals.\n\nThe dreadnought can't be harmed from within the demiplane. If\
    \ the dreadnought dies, the demiplane disappears, and everything inside it appears\
    \ around the corpse. The demiplane is otherwise indestructible."
  "name": "Demiplanar Donjon"
- "desc": "If the astral dreadnought fails a saving throw, it can choose to succeed\
    \ instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "An astral dreadnought's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "If the astral dreadnought scores a critical hit against a creature traveling\
    \ through the Astral Plane by means of the [astral projection](/3-Mechanics/CLI/spells/astral-projection.md)\
    \ spell, the dreadnought can cut the target's silver cord instead of dealing damage."
  "name": "Sever Silver Cord"
"actions":
- "desc": "The astral dreadnought makes three attacks: one with its bite and two with\
    \ its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:5d10 + 9|text(36) (5d10 + 9) piercing damage. If the\
    \ target is a creature of Huge size or smaller and this damage reduces it to 0\
    \ hit points or it is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated),\
    \ the astral dreadnought swallows it. The swallowed target, along with everything\
    \ it is wearing and carrying, appears in an unoccupied space on the floor of the\
    \ astral dreadnought's Demiplanar Donjon."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 20 ft., one\
    \ target. Hit: dice:3d6 + 9|text(19) (3d6 + 9) slashing damage."
  "name": "Claw"
"legendary_actions":
- "desc": "The astral dreadnought makes one claw attack."
  "name": "Claw"
- "desc": "One creature that is Huge or smaller that the astral dreadnought can see\
    \ within 60 feet of it must succeed on a DC 19 Charisma saving throw or be magically\
    \ teleported to an unoccupied space on the floor of the astral dreadnought's Demiplanar\
    \ Donjon. At the end of the target's next turn, the target reappears in the space\
    \ it left or in the nearest unoccupied space if that space is occupied."
  "name": "Donjon Visit (Costs 2 Actions)"
- "desc": "Each creature within 60 feet of the astral dreadnought must make a DC 19\
    \ Wisdom saving throw, taking dice:2d10 + 4|text(15) (2d10 + 4) psychic damage\
    \ on a failed save, or half as much damage on a successful one."
  "name": "Psychic Projection (Costs 3 Actions)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Astral%20Dreadnought.webp"
```
^statblock

```encounter-table
name: Astral Dreadnought
creatures:
 - 1: Astral Dreadnought
```