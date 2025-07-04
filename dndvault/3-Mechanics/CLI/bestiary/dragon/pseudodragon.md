---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-4
- monster/environment/coastal
- monster/environment/desert
- monster/environment/forest
- monster/environment/hill
- monster/environment/mountain
- monster/environment/urban
- monster/size/tiny
- monster/type/dragon
aliases: ["Pseudodragon"]
NoteIcon: monster
BestiaryType: dragon
SourceType: Bestiary
BookSource: Monster Manual p. 254. Available in the SRD.
---
# [Pseudodragon](3-Mechanics\CLI\bestiary\dragon/pseudodragon.md)
*Source: Monster Manual p. 254. Available in the SRD.*  

The elusive pseudodragon dwells in the quiet places of the world, making its home in the hollows of trees and small caves. With its red-brown scales, horns, and a maw filled with sharp teeth, a pseudodragon resembles a tiny red dragon but its disposition is playful.

## Quiet and Defensive

Pseudodragons have little interest in other creatures, and they avoid them whenever possible. If it is attacked, a pseudodragon fights back using the poisonous stinger at the tip of its tail, one jab of which can put a creature into a catatonic state that can last for hours.

## Draconic Familiars

Mages often seek out pseudodragons, whose agreeable disposition, telepathic ability, and resistance to magic make them superior familiars. Pseudodragons are selective when it comes to choosing companions, but they can sometimes be won over with gifts of food or treasure. When a pseudodragon finds an agreeable companion, it bonds with that person as long as it is treated fairly.

A pseudodragon puts up with no ill treatment, and it abandons a manipulative or abusive companion without warning.

## Language of Emotion

Pseudodragons can't speak, but they communicate using a limited form of telepathy that allows them to share basic ideas such as hunger, curiosity, or affection. When it bonds with a companion, a pseudodragon can communicate what it sees and hears even over long distances.

A pseudodragon often vocalizes animal noises. A rasping purr indicates pleasure, while a hiss means unpleasant surprise. A bird-like chirping represents desire, and a growl always means anger or discontent.

> [!note] Variant: Pseudodragon Familiar
> 
> Some pseudodragons are willing to serve spellcasters as a familiar. Such pseudodragons have the following trait.
^variant-pseudodragon-familiar

## Familiar

The pseudodragon can serve another creature as a familiar, forming a magic, telepathic bond with that willing companion. While the two are bonded, the companion can sense what the pseudodragon senses as long as they are within 1 mile of each other. While the pseudodragon is within 10 feet of its companion, the companion shares the pseudodragon's Magic Resistance trait. At any time and for any reason, the pseudodragon can end its service as a familiar, ending the telepathic bond.

```statblock
"name": "Pseudodragon"
"size": "Tiny"
"type": "dragon"
"alignment": "Neutral Good"
"ac": !!int "13"
"hp": !!int "7"
"hit_dice": "2d4 + 2"
"stats":
- !!int "6"
- !!int "15"
- !!int "13"
- !!int "10"
- !!int "12"
- !!int "10"
"speed": "15 ft., fly 60 ft."
"skillsaves":
  "Stealth": !!int "4"
  "Perception": !!int "3"
"senses": "blindsight 10 ft., darkvision 60 ft., passive Perception 13"
"languages": "understands Common and Draconic but can't speak"
"cr": "1/4"
"traits":
- "desc": "The pseudodragon has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight, hearing, or smell."
  "name": "Keen Senses"
- "desc": "The pseudodragon has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "The pseudodragon can magically communicate simple ideas, emotions, and\
    \ images telepathically with any creature within 100 feet of it that can understand\
    \ a language."
  "name": "Limited Telepathy"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing damage, and the target must\
    \ succeed on a DC 11 Constitution saving throw or become [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ for 1 hour. If the saving throw fails by 5 or more, the target falls [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ for the same duration, or until it takes damage or another creature uses an\
    \ action to shake it awake."
  "name": "Sting"
"source":
- "MM"
- "RoT"
- "ToA"
- "WDMM"
- "GoS"
- "ERLW"
- "IMR"
- "EGW"
- "IDRotF"
- "DoDk"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Pseudodragon.webp"
```
^statblock

```encounter-table
name: Pseudodragon
creatures:
 - 1: Pseudodragon
```

## Environment

mountain, forest, hill, urban, desert, coastal