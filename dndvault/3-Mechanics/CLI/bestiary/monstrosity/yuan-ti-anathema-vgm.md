---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/12
- monster/environment/desert
- monster/environment/forest
- monster/environment/underdark
- monster/size/huge
- monster/type/monstrosity/shapechanger
- monster/type/monstrosity/yuan-ti
aliases: ["Yuan-ti Anathema"]
NoteIcon: monster
BestiaryType: monstrosity (shapechanger, yuan-ti)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 202
---
# [Yuan-ti Anathema](3-Mechanics\CLI\bestiary\monstrosity/yuan-ti-anathema-vgm.md)
*Source: Volo's Guide to Monsters p. 202*  

Yuan-ti malisons who become priestly devotees of a particular god-be it Sseth, Dendar the Night Serpent, or Merrshaulk-often rise through the ranks to become spiritual leaders among the serpent folk. These priests perform sacrificial rites to appease their vile gods.

## Yuan-ti Anathema

A yuan-ti abomination's quest for godhood might lead it to perform a ritual that, if successful, transforms it into an even greater form: a yuan-ti anathema. This ritual demands the sacrifice of hundreds of snakes and requires the abomination to bathe in the blood of its enemies. The transformation is quick yet painful.

Not all yuan-ti are eager to see one of their own become an anathema, since anathemas brutally subjugate their lessers for their own evil ends.

### Not Quite Divine

An anathema considers itself a demigod on the path to greater divinity. It demands obeisance from weaker yuan-ti and uses every resource at its disposal to launch small-scale wars against its neighbors. Each conquest brings new slaves and sacrifices, as well as glory and riches, that the anathema thinks it needs to achieve true divinity.

An anathema's most loyal yuan-ti followers see it as the pinnacle of the serpentine form, an unbelievable improvement on the nearly perfect abomination. Its devoted human followers think of it as "divine flesh in a mortal body," and cultists serving an anathema tend to be more bloodthirsty and self-sacrificing in its presence.

### Immortal

Anathemas don't age, allowing them to pursue their goals until the end of days. Truly powerful ones can grow to rule multiple yuan-ti cities and bring entire regions, including humanoid realms, under yuan-ti control.

```statblock
"name": "Yuan-ti Anathema (VGM)"
"size": "Huge"
"type": "monstrosity"
"subtype": "shapechanger, yuan-ti"
"alignment": "Neutral Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "189"
"hit_dice": "18d12 + 72"
"stats":
- !!int "23"
- !!int "13"
- !!int "19"
- !!int "19"
- !!int "17"
- !!int "20"
"speed": "40 ft., climb 30 ft., swim 30 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "7"
"damage_resistances": "acid, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., blindsight 30 ft., passive Perception 17"
"languages": "Abyssal, Common, Draconic"
"cr": "12"
"traits":
- "desc": "The anathema's innate spellcasting ability is Charisma (spell save DC 17).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [animal friendship](/3-Mechanics/CLI/spells/animal-friendship.md)\
    \ (snakes only)\n\n1/day: [divine word](/3-Mechanics/CLI/spells/divine-word.md)\n\
    \n3/day each: [darkness](/3-Mechanics/CLI/spells/darkness.md), [entangle](/3-Mechanics/CLI/spells/entangle.md),\
    \ [fear](/3-Mechanics/CLI/spells/fear.md), [haste](/3-Mechanics/CLI/spells/haste.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md), [polymorph](/3-Mechanics/CLI/spells/polymorph.md)"
  "name": "Innate Spellcasting (Anathema Form Only)"
- "desc": "The anathema has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Any creature of the anathema's choice, other than a snake or a yuan-ti,\
    \ that starts its turn within 30 feet of the anathema and can see or hear it must\
    \ succeed on a DC 17 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of snakes and yuan-ti. A [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success. If a target's saving throw is successful or the\
    \ effect ends for it, the target is immune to this aura for the next 24 hours."
  "name": "Ophidiophobia Aura"
- "desc": "The anathema can use its action to polymorph into a Huge giant constrictor\
    \ snake, or back into its true form. its statistics are the same in each form.\
    \ Any equipment it is wearing or carrying isn't transformed."
  "name": "Shapechanger"
- "desc": "The anathema has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks and on saving throws against being [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded).\
    \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned),\
    \ or knocked [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)."
  "name": "Six Heads"
"actions":
- "desc": "The anathema makes two claw attacks, one constrict attack, and one Flurry\
    \ of Bites attack."
  "name": "Multiattack (Anathema Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d6 + 6|text(13) (2d6 + 6) slashing damage."
  "name": "Claw (Anathema Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 15 ft., one\
    \ Large or smaller creature. Hit: dice:3d6 + 6|text(16) (3d6 + 6) bludgeoning\
    \ damage plus dice:2d6|text(7) (2d6) acid damage, and the target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 16). Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ and takes dice:3d6 + 6|text(16) (3d6 + 6) bludgeoning damage plus dice:2d6|text(7)\
    \ (2d6) acid damage at the start of each of its turns, and the anathema can't\
    \ constrict another target."
  "name": "Constrict"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 10 ft., one\
    \ creature. Hit: dice:6d6 + 6|text(27) (6d6 + 6) piercing damage plus dice:4d6|text(14)\
    \ (4d6) poison damage."
  "name": "Flurry of Bites"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Yuan-ti%20Anathema.webp"
```
^statblock

```encounter-table
name: Yuan-ti Anathema
creatures:
 - 1: Yuan-ti Anathema
```

## Environment

underdark, forest, desert