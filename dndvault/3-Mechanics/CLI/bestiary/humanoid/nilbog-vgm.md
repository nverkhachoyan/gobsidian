---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1
- monster/environment/forest
- monster/environment/hill
- monster/environment/underdark
- monster/size/small
- monster/type/humanoid/goblinoid
aliases: ["Nilbog"]
NoteIcon: monster
BestiaryType: humanoid (goblinoid)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 182
---
# [Nilbog](3-Mechanics\CLI\bestiary\humanoid/nilbog-vgm.md)
*Source: Volo's Guide to Monsters p. 182*  

When Maglubiyet conquered the goblin gods, he intended to leave only Khurgorbaeyag alive as a harsh overseer who would keep the goblins under heel. But the goblins' pantheon included a trickster deity who was determined to get the last laugh. Although its essence was shattered by Maglubiyet, this trickster god survives in splintered form as a possessing spirit that arises when goblinoids form a host, causing disorder in the ranks unless it is appeased. Goblins have no name for this deity and dare not give it one, lest Maglubiyet use its name to ensnare and crush it as he did their other deities. They call the possessing spirit, and the goblin possessed by it, a nilbog ("goblin" spelled backward), and they revel in the fear that a nilbog sows among the ranks of the bugbears and hobgoblins in the host.

## Goblins' Revenge

When goblinoids form a host, there is a chance that a goblin will become possessed by a nilbog, particularly if the goblin has been mistreated by its betters. This possession turns the goblin into a wisecracking, impish creature fearless of reprisal. It gives the goblin strange powers that drive others to do the opposite of what they desire. Attacking a goblin possessed by a nilbog is foolhardy, and killing the creature just prompts the spirit to possess another goblin. The only way to keep a nilbog from wreaking havoc is to treat it well and give it respect and praise.

## No Joking Matter

The possible presence of a nilbog in a host has given rise to a practice among goblinoids that each host include at least one goblin jester. This jester is allowed to go anywhere and do whatever it pleases. The position of jester is a much sought-after one among the goblins, because even if the jester is obviously not a nilbog, hobgoblins and bugbears indulge its manic behavior.

## Nilbogism

A nilbog is an invisible spirit that possesses only goblins. Bereft of a host, the spirit has a flying speed of 30 feet and can't speak or be attacked. The only action it can take is to attempt to possess a goblin within 5 feet of it.

A goblin targeted by the spirit must succeed on a DC 15 Charisma saving throw or become possessed. While possessed by the spirit, the goblin's alignment becomes chaotic evil, its Charisma becomes 15 (unless it was already higher), and it gains the nilbog's Innate Spellcasting and Nilbogism traits, as well as its Reversal of Fortune reaction. If the save succeeds, the spirit can't possess that goblin for 24 hours. If its host is killed or the possession is ended by a spell such as hallow, magic circle, or protection from evil and good, the spirit searches for another goblin to possess. The spirit can leave its host at any time, but it won't do so willingly unless it knows there's another potential host nearby. A goblin stripped of its nilbog spirit reverts to its normal statistics and loses the traits it gained while possessed.

```statblock
"name": "Nilbog (VGM)"
"size": "Small"
"type": "humanoid"
"subtype": "goblinoid"
"alignment": "Chaotic Evil"
"ac": !!int "13"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "7"
"hit_dice": "2d6"
"stats":
- !!int "8"
- !!int "14"
- !!int "10"
- !!int "10"
- !!int "8"
- !!int "15"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "6"
"senses": "darkvision 60 ft., passive Perception 9"
"languages": "Common, Goblin"
"cr": "1"
"traits":
- "desc": "The nilbog's innate spellcasting ability is Charisma (spell save DC 12).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [Tasha's hideous\
    \ laughter](/3-Mechanics/CLI/spells/tashas-hideous-laughter.md), [vicious mockery](/3-Mechanics/CLI/spells/vicious-mockery.md)\n\
    \n1/day: [confusion](/3-Mechanics/CLI/spells/confusion.md)"
  "name": "Innate Spellcasting"
- "desc": "Any creature that attempts to damage the nilbog must first succeed on a\
    \ DC 12 Charisma saving throw or be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ until the end of the creature's next turn. A creature [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ in this way must use its action praising the nilbog. The nilbog can't regain\
    \ hit points, including through magical healing, except through its Reversal of\
    \ Fortune reaction."
  "name": "Nilbogism"
- "desc": "The nilbog can take the Disengage or Hide action as a bonus action on each\
    \ of its turns."
  "name": "Nimble Escape"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) bludgeoning damage."
  "name": "Fool's Scepter"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Shortbow"
"reactions":
- "desc": "In response to another creature dealing damage to the nilbog, the nilbog\
    \ reduces the damage to 0 and regains dice: 1d6|avg|noform (1d6) hit points."
  "name": "Reversal of Fortune"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Nilbog.webp"
```
^statblock

```encounter-table
name: Nilbog
creatures:
 - 1: Nilbog
```

## Environment

underdark, forest, hill