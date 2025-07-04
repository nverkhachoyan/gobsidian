---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/4
- monster/environment/desert
- monster/environment/forest
- monster/environment/underdark
- monster/size/medium
- monster/type/monstrosity/shapechanger
- monster/type/monstrosity/yuan-ti
aliases: ["Yuan-ti Mind Whisperer"]
NoteIcon: monster
BestiaryType: monstrosity (shapechanger, yuan-ti)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 204
---
# [Yuan-ti Mind Whisperer](3-Mechanics\CLI\bestiary\monstrosity/yuan-ti-mind-whisperer-vgm.md)
*Source: Volo's Guide to Monsters p. 204*  

Yuan-ti malisons who become priestly devotees of a particular god-be it Sseth, Dendar the Night Serpent, or Merrshaulk-often rise through the ranks to become spiritual leaders among the serpent folk. These priests perform sacrificial rites to appease their vile gods.

## Yuan-ti Mind Whisperer

Mind whisperers are malison spellcasters that enter into a pact with the serpent god Sseth, the Sibilant Death. They use their abilities to convert others to their faith, increase their personal power, and befuddle the minds of their enemies.

A mind whisperer is elusive, manipulative, unpredictable, and willing to cheat or kill comrades and rivals alike if doing so benefits it. The worshipers of Sseth have their hands in many schemes, often plying the middle ground between two factions, and thus spend a lot of energy making sure neither of their allies learn of their conflicting connections. Even among yuan-ti, mind whisperers are known for being self-important, sneaky, and prone to flee at the first sign of trouble.

This malison is the type that has a human body and a snake head.

```statblock
"name": "Yuan-ti Mind Whisperer (VGM)"
"size": "Medium"
"type": "monstrosity"
"subtype": "shapechanger, yuan-ti"
"alignment": "Neutral Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "71"
"hit_dice": "13d8 + 13"
"stats":
- !!int "16"
- !!int "14"
- !!int "13"
- !!int "14"
- !!int "14"
- !!int "16"
"speed": "30 ft."
"saves":
  "Charisma": !!int "5"
  "Wisdom": !!int "4"
"skillsaves":
  "Deception": !!int "5"
  "Stealth": !!int "4"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft. (penetrates magical darkness), passive Perception 12"
"languages": "Abyssal, Common, Draconic"
"cr": "4"
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
    \ [message](/3-Mechanics/CLI/spells/message.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [poison spray](/3-Mechanics/CLI/spells/poison-spray.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\
    \n1st-3rd level (2 slots): [charm person](/3-Mechanics/CLI/spells/charm-person.md),\
    \ [crown of madness](/3-Mechanics/CLI/spells/crown-of-madness.md), [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md),\
    \ [expeditious retreat](/3-Mechanics/CLI/spells/expeditious-retreat.md), [fly](/3-Mechanics/CLI/spells/fly.md),\
    \ [hypnotic pattern](/3-Mechanics/CLI/spells/hypnotic-pattern.md), [illusory script](/3-Mechanics/CLI/spells/illusory-script.md)"
  "name": "Spellcasting (Yuan-ti Form Only)"
- "desc": "The yuan-ti can use its action to polymorph into a Medium snake or back\
    \ into its true form. Its statistics are the same in each form. Any equipment\
    \ it is wearing or carrying isn't transformed. If it dies, it stays in its current\
    \ form."
  "name": "Shapechanger"
- "desc": "The yuan-ti has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The first time the yuan-ti hits with a melee attack on its turn, it can\
    \ deal an extra dice:3d10|text(16) (3d10) psychic damage to the target."
  "name": "Mind Fangs (2/Day)"
- "desc": "When the yuan-ti reduces an enemy to 0 hit points, the yuan-ti gains 9\
    \ temporary hit points."
  "name": "Sseth's Blessing"
"actions":
- "desc": "The yuan-ti makes one bite attack and one scimitar attack."
  "name": "Multiattack (Yuan-ti Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage plus dice:2d6|text(7)\
    \ (2d6) poison damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) slashing damage."
  "name": "Scimitar (Yuan-ti Form Only)"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Yuan-ti%20Mind%20Whisperer.webp"
```
^statblock

```encounter-table
name: Yuan-ti Mind Whisperer
creatures:
 - 1: Yuan-ti Mind Whisperer
```

## Environment

underdark, forest, desert