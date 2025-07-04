---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/undead
aliases: ["Ghoul"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Monster Manual p. 148, Divine Contention, Storm Lord's Wrath, Sleeping Dragon's Wake, Tasha's Cauldron of Everything, Dragon of Icespire Peak. Available in the SRD and the Basic Rules.
---
# [Ghoul](3-Mechanics\CLI\bestiary\undead/ghoul.md)
*Source: Monster Manual p. 148, Divine Contention, Storm Lord's Wrath, Sleeping Dragon's Wake, Tasha's Cauldron of Everything, Dragon of Icespire Peak. Available in the SRD and the Basic Rules.*  

Ghouls roam the night in packs, driven by an insatiable hunger for humanoid flesh.

## Devourers of Flesh

Like maggots or carrion beetles, ghouls thrive in places rank with decay and death. A ghoul haunts a place where it can gorge on dead flesh and decomposing organs. When it can't feed on the dead, it pursues living creatures and attempts to make corpses of them. Though they gain no nourishment from the corpses they devour, ghouls are driven by an unending hunger that compels them to consume. A ghoul's undead flesh never rots, and this monster can persist in a crypt or tomb for untold ages without feeding.

## Abyssal Origins

Ghouls trace their origins to the Abyss. Doresain, the first of their kind, was an elf worshiper of Orcus. Turning against his own people, he feasted on humanoid flesh to honor the Demon Prince of Undeath. As a reward for his service, Orcus transformed Doresain into the first ghoul. Doresain served Orcus faithfully in the Abyss, creating ghouls from the demon lord's other servants until an incursion by Yeenoghu, the demonic Gnoll Lord, robbed Doresain of his abyssal domain. When Orcus would not intervene on his behalf, Doresain turned to the elf gods for salvation, and they took pity on him and helped him escape certain destruction. Since then, elves have been immune to the ghouls' paralytic touch.

## Ghasts

Orcus sometimes infuses a ghoul with a stronger dose of abyssal energy, making a ghast. Whereas ghouls are little more than savage beasts, a ghast is cunning and can inspire a pack of ghouls to follow its commands.

```statblock
"name": "Ghoul"
"size": "Medium"
"type": "undead"
"alignment": "Chaotic Evil"
"ac": !!int "12"
"hp": !!int "22"
"hit_dice": "5d8"
"stats":
- !!int "13"
- !!int "15"
- !!int "10"
- !!int "7"
- !!int "10"
- !!int "6"
"speed": "30 ft."
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Common"
"cr": "1"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+2 (+2) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 2|text(7) (2d4 + 2) slashing damage. If the target is\
    \ a creature other than an elf or undead, it must succeed on a DC 10 Constitution\
    \ saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for 1 minute. The target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success."
  "name": "Claws"
"source":
- "MM"
- "CoS"
- "LMoP"
- "PotA"
- "TftYP"
- "ToA"
- "WDMM"
- "GoS"
- "DC"
- "SLW"
- "SDW"
- "BGDIA"
- "RMBRE"
- "IDRotF"
- "TCE"
- "CM"
- "WBtW"
- "JttRC"
- "DoSI"
- "KftGV"
- "PSI"
- "HftT"
- "PaBTSO"
- "DIP"
- "AATM"
- "SatO"
- "BMT"
- "GHLoE"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Ghoul.webp"
```
^statblock

```encounter-table
name: Ghoul
creatures:
 - 1: Ghoul
```

## Environment

underdark, swamp, urban