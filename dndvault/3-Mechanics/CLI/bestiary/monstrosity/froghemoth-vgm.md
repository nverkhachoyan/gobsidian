---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/10
- monster/environment/swamp
- monster/environment/underdark
- monster/size/huge
- monster/type/monstrosity
aliases: ["Froghemoth"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 145
---
# [Froghemoth](3-Mechanics\CLI\bestiary\monstrosity/froghemoth-vgm.md)
*Source: Volo's Guide to Monsters p. 145*  

A froghemoth is an amphibious predator as big as an elephant. It lairs in swamps and has four tentacles, a thick rubbery hide, a fang-filled maw with a prehensile tongue, and an extendable stalk sporting three bulbous eyes that face in different directions.

## Otherworldly Entities

Froghemoths are creatures not of this world. A journal purportedly written long ago by the wizard Lum the Mad describes strange, cylindrical chambers of metal buried in the ground from which froghemoths emerged, but no reliable reports of the location of such places exist.

## Hungry from Birth

Every few years, a froghemoth can lay a fertile egg without mating. The froghemoth cares nothing for its egg, and might eat the hatchling. A young froghemoth's survival is most often predicated on its parent leaving it behind in indifference. A newborn froghemoth grows to full size over a period of months by indiscriminately preying on other creatures in its swampy domain. It learns to hide its enormous body in murky pools, keeping only its eyestalk above water to watch for passing creatures. When food comes within reach, the froghemoth erupts from its pool, tentacles and tongue flailing. It can grab several targets at once, keeping them at bay while it wraps its tongue around another one and pulls it in to be devoured.

## Revered by Bullywugs

If a bullywug tribe comes across a froghemoth, the bullywugs treat the froghemoth as a god and do all they can to coax the monster into their den. A froghemoth can be tamed (after a fashion) by offering it food, and bullywugs can communicate with it on a basic level, so the creature might eat only a few bullywugs before following the rest. Bullywugs gather food as tribute for it, provide it with a comfortable lair, fanatically protect it from harm, and try to ensure that any young froghemoth reaches maturity.

```statblock
"name": "Froghemoth (VGM)"
"size": "Huge"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "184"
"hit_dice": "16d12 + 80"
"stats":
- !!int "23"
- !!int "13"
- !!int "20"
- !!int "2"
- !!int "12"
- !!int "5"
"speed": "30 ft., swim 30 ft."
"saves":
  "Wisdom": !!int "5"
  "Constitution": !!int "9"
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "9"
"damage_resistances": "fire, lightning"
"senses": "darkvision 60 ft., passive Perception 19"
"languages": ""
"cr": "10"
"traits":
- "desc": "The froghemoth can breathe air and water."
  "name": "Amphibious"
- "desc": "If the froghemoth takes lightning damage, it suffers several effects until\
    \ the end of its next turn: its speed is halved, it takes a −2 penalty to AC and\
    \ Dexterity saving throws, it can't use reactions or Multiattack, and on its turn,\
    \ it can use either an action or a bonus action, not both."
  "name": "Shock Susceptibility"
"actions":
- "desc": "The froghemoth makes two attacks with its tentacles. It can also use its\
    \ tongue or bite."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 20 ft., one\
    \ target. Hit: dice:3d8 + 6|text(19) (3d8 + 6) bludgeoning damage, and the\
    \ target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape\
    \ DC 16) if it is a Huge or smaller creature. Until the grapple ends, the froghemoth\
    \ can't use this tentacle on another target. The froghemoth has four tentacles."
  "name": "Tentacle"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:3d10 + 6|text(22) (3d10 + 6) piercing damage, and the\
    \ target is swallowed if it is a Medium or smaller creature. A swallowed creature\
    \ is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) and [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ has total cover against attacks and other effects outside the froghemoth, and\
    \ takes dice:3d6|text(10) (3d6) acid damage at the start of each of the froghemoth's\
    \ turns.\n\nThe froghemoth's gullet can hold up to two creatures at a time. If\
    \ the Froghemoth takes 20 damage or more on a single turn from a creature inside\
    \ it, the Froghemoth must succeed on a DC 20 Constitution saving throw at the\
    \ end of that turn or regurgitate all swallowed creatures, each of which falls\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone) in a space within 10 feet\
    \ of the froghemoth. If the froghemoth dies, a swallowed creature is no longer\
    \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained) by it and can\
    \ escape from the corpse using 10 feet of movement, exiting [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Bite"
- "desc": "The Froghemoth targets one Medium or smaller creature that it can see within\
    \ 20 feet of it. The target must make a DC 18 Strength saving throw. On a failed\
    \ save, the target is pulled into an unoccupied space within 5 feet of the froghemoth,\
    \ and the froghemoth can make a bite attack against it as a bonus action."
  "name": "Tongue"
"source":
- "VGM"
- "ToA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Froghemoth.webp"
```
^statblock

```encounter-table
name: Froghemoth
creatures:
 - 1: Froghemoth
```

## Environment

underdark, swamp