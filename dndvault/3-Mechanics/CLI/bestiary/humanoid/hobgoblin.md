---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-2
- monster/environment/desert
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/goblinoid
aliases: ["Hobgoblin"]
NoteIcon: monster
BestiaryType: humanoid (goblinoid)
SourceType: Bestiary
BookSource: Monster Manual p. 186. Available in the SRD and the Basic Rules.
---
# [Hobgoblin](3-Mechanics\CLI\bestiary\humanoid/hobgoblin.md)
*Source: Monster Manual p. 186. Available in the SRD and the Basic Rules.*  

War horns sound, stones fly from catapults, and the thunder of a thousand booted feet echoes across the land as hobgoblins march to battle. Across the borderlands of civilization, settlements and settlers must contend with these aggressive humanoids, whose thirst for conquest is never satisfied.

Hobgoblins have dark orange or red-orange skin, and hair ranging from dark red-brown to dark gray. Yellow or dark brown eyes peer out beneath their beetling brows, and their wide mouths sport sharp and yellowed teeth. A male hobgoblin might have a large blue or red nose, which symbolizes virility and power among goblinkin. Hobgoblins can live as long as humans, though their love of warfare and battle means that few do.

## Goblinoids

Hobgoblins belong to a family of creatures called goblinoids. They are often found lording over their cousins, the smaller goblins and the ferocious bugbears.

## Martial Might

A hobgoblin measures virtue by physical strength and martial prowess, caring about nothing except the opportunity to demonstrate skill and cunning in battle. Hobgoblins of high military rank attain their positions by force, then hold those positions by imposing their authority through draconian measures.

Hobgoblins train to fight with a variety of weapons, and have great skill at crafting arms, armor, siege engines, and other military devices. Organized and disciplined, they take exceptional care of their weapons, armor, and personal possessions. They favor the bold colors associated with their tribes, and trim their often-elaborate uniforms with blood-red piping and leather dyed black.

## Military Legions

Hobgoblins organize themselves into tribal bands known as legions. In their martial society, every hobgoblin has a rank, from the powerful leaders and champions, to the rank-and-file foot soldiers, to the goblins that find themselves driven into the front lines at spear point. A legion is headed by a warlord with several captains serving under its command. A hobgoblin warlord is a ruthless tyrant more interested in strategy, victory, glory, reputation, and dominion than leading troops into battle.

As loyal and disciplined as hobgoblins are in their own legion, rival legions compete constantly for reputation and status. Meetings between legions erupt in violence if troops aren't restrained, and only exceptionally powerful leaders can force legions to cooperate on the battlefield.

## Strategic Thinkers

Hobgoblins have a strong grasp of tactics and discipline, and can carry out sophisticated battle plans under the direction of a strategically minded leader. However, they hate elves and attack them first in battle over any other opponents, even if doing so would be a tactical error.

Legions often supplement their ranks with less reliable and more expendable troops, including goblins, bugbears, orcs, evil humans, ogres, and giants.

## Beast Trainers

Hobgoblins have a long history of training animals to service. Like the more civilized races, they use oxen and horses to transport goods and weaponry over long distances. They communicate with each other using trained ravens, and keep vicious wolves to guard prisoners and protect hobgoblin camps. Hobgoblin cavalry use trained worgs as steeds, in the same way that goblins ride wolves. Some tribes even keep carnivorous apes as fighting beasts.

## Conquer and Control

Hobgoblins claim lands with abundant resources, and they can be found in forests and mountains, near mines and humanoid settlements, and anywhere else that wood, metal, and potential slaves can be found. They build and conquer strongholds in strategically advantageous locations, which they then use as staging areas to expand their territory.

Hobgoblin warlords never tire of combat, but they don't take up arms lightly. Before they attack, hobgoblins conduct thorough reconnaissance to gauge the strengths and weaknesses of their foes. When assaulting a stronghold, they surround it first to cut off escape routes and supply lines, then slowly starve their enemies out.

Hobgoblins fortify their own holdings, bolstering existing defenses with innovations of their own. Whether they lair in cavern complexes, dungeons, ruins, or forests, they protect their strongholds with ditches, fences, gates, guard towers, pit traps, and crude catapults or ballistas.

## Legion of Maglubiyet

Hobgoblins worship Maglubiyet the Mighty One, the greater god of goblinoids. As terrifying as this figure is, hobgoblins don't fear death, believing that when they die in battle, their spirits join the honored ranks of Maglubiyet's army on the plane of Acheron.

> [!quote]- A quote from Translation of a Hobgoblin War Chant  
> 
> They break before our shields,
> 
> They fall beneath our blades;
> 
> Their home is ours to conquer,
> 
> Their children our slaves.
> 
> Acheron! Acheron!
> 
> Victory is ours!


```statblock
"name": "Hobgoblin"
"size": "Medium"
"type": "humanoid"
"subtype": "goblinoid"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "[chain mail](/3-Mechanics/CLI/items/chain-mail.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "11"
"hit_dice": "2d8 + 2"
"stats":
- !!int "13"
- !!int "12"
- !!int "12"
- !!int "10"
- !!int "10"
- !!int "9"
"speed": "30 ft."
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Common, Goblin"
"cr": "1/2"
"traits":
- "desc": "Once per turn, the hobgoblin can deal an extra dice:2d6|text(7) (2d6)\
    \ damage to a creature it hits with a weapon attack if that creature is within\
    \ 5 feet of an ally of the hobgoblin that isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Martial Advantage"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 1|text(5) (1d8 + 1) slashing damage, or dice:1d10 + 1|text(6)\
    \ (1d10 + 1) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+3 (+3) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 1|text(5) (1d8 + 1) piercing damage."
  "name": "Longbow"
"source":
- "MM"
- "HotDQ"
- "LMoP"
- "PotA"
- "SKT"
- "TftYP"
- "WDMM"
- "GoS"
- "BGDIA"
- "ERLW"
- "RMBRE"
- "EGW"
- "WBtW"
- "DSotDQ"
- "PaBTSO"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Hobgoblin.webp"
```
^statblock

```encounter-table
name: Hobgoblin
creatures:
 - 1: Hobgoblin
```

## Environment

underdark, grassland, forest, hill, desert