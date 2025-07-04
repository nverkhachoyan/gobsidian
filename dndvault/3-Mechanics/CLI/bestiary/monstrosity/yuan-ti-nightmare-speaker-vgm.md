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
aliases: ["Yuan-ti Nightmare Speaker"]
NoteIcon: monster
BestiaryType: monstrosity (shapechanger, yuan-ti)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 205
---
# [Yuan-ti Nightmare Speaker](3-Mechanics\CLI\bestiary\monstrosity/yuan-ti-nightmare-speaker-vgm.md)
*Source: Volo's Guide to Monsters p. 205*  

Yuan-ti malisons who become priestly devotees of a particular god-be it Sseth, Dendar the Night Serpent, or Merrshaulk-often rise through the ranks to become spiritual leaders among the serpent folk. These priests perform sacrificial rites to appease their vile gods.

## Yuan-ti Nightmare Speaker

Nightmare speakers are female yuan-ti malison priests that make a pact with the Dendar the Night Serpent to feed their deity the fears and nightmares of their victims in exchange for power in the mortal world. The priestesses receive nightmarish visions from Dendar, which they interpret as prophecies, and then use their magic and influence to make these visions come true.

The cruelest of all yuan-ti, nightmare speakers revel in torturing prisoners and slaves, leaving them in a constant state of fear and dread. They prefer to terrify rather than kill their opponents. They manipulate humanoid communities for the purpose of acquiring more victims, and enjoy the company of undead.

This malison is the type that has a human head and upper body with a serpentine lower body instead of legs.

```statblock
"name": "Yuan-ti Nightmare Speaker (VGM)"
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
- !!int "12"
- !!int "16"
"speed": "30 ft."
"saves":
  "Charisma": !!int "5"
  "Wisdom": !!int "3"
"skillsaves":
  "Deception": !!int "5"
  "Stealth": !!int "4"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft. (penetrates magical darkness), passive Perception 11"
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
    \ following warlock spells:\n\nCantrips (at will): [chill touch](/3-Mechanics/CLI/spells/chill-touch.md),\
    \ [eldritch blast](/3-Mechanics/CLI/spells/eldritch-blast.md) (range 300 ft.,\
    \ +3 bonus to each damage roll), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [message](/3-Mechanics/CLI/spells/message.md), [poison spray](/3-Mechanics/CLI/spells/poison-spray.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\n1st-3rd\
    \ level (2 slots): [arms of Hadar](/3-Mechanics/CLI/spells/arms-of-hadar.md),\
    \ [darkness](/3-Mechanics/CLI/spells/darkness.md), [fear](/3-Mechanics/CLI/spells/fear.md),\
    \ [hex](/3-Mechanics/CLI/spells/hex.md), [hold person](/3-Mechanics/CLI/spells/hold-person.md),\
    \ [hunger of Hadar](/3-Mechanics/CLI/spells/hunger-of-hadar.md), [witch bolt](/3-Mechanics/CLI/spells/witch-bolt.md)"
  "name": "Spellcasting (Yuan-ti Form Only)"
- "desc": "The yuan-ti can use its action to polymorph into a Medium snake or back\
    \ into its true form. Its statistics are the same in each form. Any equipment\
    \ it is wearing or carrying isn't transformed. If it dies, it stays in its current\
    \ form."
  "name": "Shapechanger"
- "desc": "The first time the yuan-ti hits with a melee attack on its turn, it can\
    \ deal an extra dice:3d10|text(16) (3d10) necrotic damage to the target."
  "name": "Death Fangs (2/Day)"
- "desc": "The yuan-ti has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The yuan-ti makes one constrict attack and one scimitar attack."
  "name": "Multiattack (Yuan-ti Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) bludgeoning damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 14)\
    \ if it is a Large or smaller creature. Until this grapple ends, the target is\
    \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained), and the yuan-ti\
    \ can't constrict another target."
  "name": "Constrict"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) slashing damage."
  "name": "Scimitar (Yuan-ti Form Only)"
- "desc": "The yuan-ti taps into the nightmares of a creature it can see within 60\
    \ feet of it and creates an illusory, immobile manifestation of the creature's\
    \ deepest fears, visible only to that creature. The target must make a DC 13 Intelligence\
    \ saving throw. On a failed save, the target takes dice:2d10|text(11) (2d10)\
    \ psychic damage and is [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of the manifestation, believing it to be real. The yuan-ti must concentrate\
    \ to maintain the illusion (as if concentrating on a spell), which lasts for up\
    \ to 1 minute and can't be harmed. The target can repeat the saving throw at the\
    \ end of each of its turns, ending the illusion on a success, or taking dice:2d10|text(11)\
    \ (2d10) psychic damage on a failure."
  "name": "Invoke Nightmare (Recharges after a Short or Long Rest)"
"source":
- "VGM"
- "ToA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Yuan-ti%20Nightmare%20Speaker.webp"
```
^statblock

```encounter-table
name: Yuan-ti Nightmare Speaker
creatures:
 - 1: Yuan-ti Nightmare Speaker
```

## Environment

underdark, forest, desert