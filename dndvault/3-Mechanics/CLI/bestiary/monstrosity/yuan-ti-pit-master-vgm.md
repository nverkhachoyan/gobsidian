---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/desert
- monster/environment/forest
- monster/environment/underdark
- monster/size/medium
- monster/type/monstrosity/shapechanger
- monster/type/monstrosity/yuan-ti
aliases: ["Yuan-ti Pit Master"]
NoteIcon: monster
BestiaryType: monstrosity (shapechanger, yuan-ti)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 206
---
# [Yuan-ti Pit Master](3-Mechanics\CLI\bestiary\monstrosity/yuan-ti-pit-master-vgm.md)
*Source: Volo's Guide to Monsters p. 206*  

Yuan-ti malisons who become priestly devotees of a particular god-be it Sseth, Dendar the Night Serpent, or Merrshaulk-often rise through the ranks to become spiritual leaders among the serpent folk. These priests perform sacrificial rites to appease their vile gods.

## Yuan-ti Pit Master

Pit masters are yuan-ti malison priests that have made a pact with the god Merrshaulk and seek to rouse him from his slumber by sacrificing humanoids to him. They are the most traditionalist in attitude among yuan-ti and believe that they are best equipped to achieve the goals of their people.

Pit masters are deeply involved in the race's long-term plan to take over humanoid governments, as well as in the ongoing effort to protect their cities from discovery or attacks by hostiles. They oppose reckless behavior and argue for a slow, cautious approach in all matters.

This malison is the type that has a human head and body and snakes for arms.

```statblock
"name": "Yuan-ti Pit Master (VGM)"
"size": "Medium"
"type": "monstrosity"
"subtype": "shapechanger, yuan-ti"
"alignment": "Neutral Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "88"
"hit_dice": "16d8 + 16"
"stats":
- !!int "16"
- !!int "14"
- !!int "13"
- !!int "14"
- !!int "12"
- !!int "16"
"speed": "30 ft."
"saves":
  "Charisma": !!int "6"
  "Wisdom": !!int "4"
"skillsaves":
  "Deception": !!int "6"
  "Stealth": !!int "5"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft. (penetrates magical darkness), passive Perception 11"
"languages": "Abyssal, Common, Draconic"
"cr": "5"
"traits":
- "desc": "The yuan-ti's innate spellcasting ability is Charisma (spell save DC 13).\
    \ The yuan-ti can innately cast the following spells, requiring no material components:\n\
    \nAt will: [animal friendship](/3-Mechanics/CLI/spells/animal-friendship.md)\
    \ (snakes only)\n\n3/day: [suggestion](/3-Mechanics/CLI/spells/suggestion.md)"
  "name": "Innate Spellcasting (Yuan-ti Form Only)"
- "desc": "The yuan-ti is a 6th-level spellcaster. Its spellcasting ability is Charisma\
    \ (spell save DC 13, dice: d20+5 (+5) to hit with spell attacks). It regains\
    \ its expended spell slots when it finishes a short or long rest. It knows the\
    \ following warlock spells:\n\nCantrips (at will): [eldritch blast](/3-Mechanics/CLI/spells/eldritch-blast.md)\
    \ (range 300 ft., +3 bonus to each damage roll), [friends](/3-Mechanics/CLI/spells/friends.md),\
    \ [guidance](/3-Mechanics/CLI/spells/guidance.md), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [message](/3-Mechanics/CLI/spells/message.md), [poison spray](/3-Mechanics/CLI/spells/poison-spray.md)\n\
    \n1st-3rd level (2 slots): [command](/3-Mechanics/CLI/spells/command.md),\
    \ [counterspell](/3-Mechanics/CLI/spells/counterspell.md), [hellish rebuke](/3-Mechanics/CLI/spells/hellish-rebuke.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md), [misty step](/3-Mechanics/CLI/spells/misty-step.md),\
    \ [unseen servant](/3-Mechanics/CLI/spells/unseen-servant.md), [vampiric touch](/3-Mechanics/CLI/spells/vampiric-touch.md)"
  "name": "Spellcasting (Yuan-ti Form Only)"
- "desc": "The yuan-ti can use its action to polymorph into a Medium snake or back\
    \ into its true form. Its statistics are the same in each form. Any equipment\
    \ it is wearing or carrying isn't transformed. It doesn't change form if it dies."
  "name": "Shapechanger"
- "desc": "The yuan-ti has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The first time the yuan-ti hits with a melee attack on its turn, it can\
    \ deal an extra dice:3d10|text(16) (3d10) poison damage to the target."
  "name": "Poison's Disciple (2/Day)"
"actions":
- "desc": "The yuan-ti makes two bite attacks using its snake arms."
  "name": "Multiattack (Yuan-ti Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage plus dice:2d6|text(7)\
    \ (2d6) poison damage."
  "name": "Bite"
- "desc": "The yuan-ti targets up to five creatures that it can see within 60 feet\
    \ of it. Each target must succeed on a DC 13 Constitution saving throw or fall\
    \ into a magical sleep and be [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ for 10 minutes. A sleeping target awakens if it takes damage or if someone uses\
    \ an action to shake or slap it awake. This magical sleep has no effect on a creature\
    \ immune to being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)."
  "name": "Merrshaulk's Slumber (1/Day)"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Yuan-ti%20Pit%20Master.webp"
```
^statblock

```encounter-table
name: Yuan-ti Pit Master
creatures:
 - 1: Yuan-ti Pit Master
```

## Environment

underdark, forest, desert