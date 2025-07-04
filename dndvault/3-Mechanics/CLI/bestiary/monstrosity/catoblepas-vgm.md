---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/swamp
- monster/size/large
- monster/type/monstrosity
aliases: ["Catoblepas"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 129
---
# [Catoblepas](3-Mechanics\CLI\bestiary\monstrosity/catoblepas-vgm.md)
*Source: Volo's Guide to Monsters p. 129*  

The catoblepas is as loathsome as the vile swamplands in which it lives. Like such wastelands, this conglomeration of bloated buffalo, dinosaur, warthog, and hippopotamus parts has few redeeming qualities. Few travelers willingly traverse the territory of a catoblepas.

## Animalistic Nature

Despite their ungainly physiology, catoblepases resemble natural beasts. A catoblepas behaves much like an animal, too, ambling through its marshy home, munching choice vegetation, eating the occasional bit of carrion, and wallowing in mire. A catoblepas might be found with the one mate it chooses for life and, on occasion, a calf. Especially if it's guarding its young, a catoblepas attacks anyone that moves too close.

Sages say that gods of pestilence and rot created catoblepases as embodiments of their influence. Whatever the origin of the creature, stories link the catoblepas to misfortune, and many of these yarns have elements of truth. Some such tales claim that hags tend catoblepases like cattle, and that a swamp that contains a catoblepas might also be home to a hag that drinks the monster's milk. Although a particular catoblepas might not be linked to a hag, a coven of hags might keep one or more of these beasts as guardians or pets. Other legends say that those of impure heart can tame a catoblepas. Indeed, some tales have circulated of malevolent warlocks and dark knights who have discovered how to domesticate the beasts and use them as mounts.

## Stench of Death

A catoblepas's stink, like that of death mixed with swamp gas and skunk musk, gives it away as being much more ghastly than its appearance suggests. When it is on the attack, a catoblepas reveals the extent of its horrific nature. The creature's serpentine neck has trouble lifting its head, but one glare from its bloodshot eyes can rot flesh. At the end of its tail is a club that can rattle body and soul if it strikes true, leaving a victim unable to act. If the target of its attacks dies, the catoblepas feasts on the fresh remains.

## Blighted Territory

A catoblepas's nature as a creature of disease and decay brings out similar characteristics in the creature's swampy habitat. Such a wetland becomes gloomy, tangled, and more fetid than it was before. Beneficial qualities of the environment, such as healing herbs and clean water, diminish when a catoblepas lives nearby. Swamp gases have a hint of the catoblepas's foulness to them. Animals in the area are more aggressive and liable to be diseased. Degenerate creatures are likely to take up residence near a catoblepas's territory, as are those seeking to avoid notice.

## Sinister Folklore

Ordinary folk rarely see a catoblepas, but the creature has such a feared reputation that stories about it are ingrained in the popular culture. Any rumor of a catoblepas taking up residence nearby is taken to be a bad omen, even if the rumor is proven false. The silhouette of a catoblepas, with its tail extended over its body and its head held low, is a baleful heraldic figure signifying death or doom.

```statblock
"name": "Catoblepas (VGM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "84"
"hit_dice": "8d10 + 40"
"stats":
- !!int "19"
- !!int "12"
- !!int "21"
- !!int "3"
- !!int "14"
- !!int "8"
"speed": "30 ft."
"senses": "darkvision 60 ft., passive Perception 12"
"languages": ""
"cr": "5"
"traits":
- "desc": "The catoblepas has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
- "desc": "Any creature other than a catoblepas that starts its turn within 10 feet\
    \ of the catoblepas must succeed on a DC 16 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until the start of the creature's next turn. On a successful saving throw, the\
    \ creature is immune to the stench of any catoblepas for 1 hour."
  "name": "Stench"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:5d6 + 4|text(21) (5d6 + 4) bludgeoning damage, and the target\
    \ must succeed on a DC 16 Constitution saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the start of the catoblepas's next turn."
  "name": "Tail"
- "desc": "The catoblepas targets a creature that it can see within 30 feet of it.\
    \ The target must make a DC 16 Constitution saving throw, taking dice:8d8|text(36)\
    \ (8d8) necrotic damage on a failed save, or half as much damage on a successful\
    \ one. If the saving throw fails by 5 or more, the target instead takes 64 necrotic\
    \ damage. The target dies if reduced to 0 hit points by this ray."
  "name": "Death Ray (Recharge 5-6)"
"source":
- "VGM"
- "MOT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Catoblepas.webp"
```
^statblock

```encounter-table
name: Catoblepas
creatures:
 - 1: Catoblepas
```

## Environment

swamp