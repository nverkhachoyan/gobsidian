---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-4
- monster/environment/forest
- monster/size/tiny
- monster/type/fey
aliases: ["Pixie"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Monster Manual p. 253
---
# [Pixie](3-Mechanics\CLI\bestiary\fey/pixie.md)
*Source: Monster Manual p. 253*  

Standing barely a foot tall, pixies resemble diminutive elves with gossamer wings like those of dragonflies or butterflies, bright as the clear dawn and as luminous as the full moonrise.

Curious as cats and shy as deer, pixies go where they please. They like to spy on other creatures and can barely contain their excitement around them. The urge to introduce themselves and strike up a friendship is almost overwhelming; only a pixie's fear of being captured or attacked stays its hand. Those who wander through a pixie's glade might never see the creatures, yet hear the occasional giggle, gasp, or sigh.

Pixies array themselves like princes and princesses of the fey, wearing flowing gowns and doublets of silk that sparkle like moonlight on a pond. Some dress in acorns, leaves, bark, and the pelts of tiny woodland beasts. They take great pride in their regalia and beam with joy when they are complimented on their ensembles.

## Magical Faerie Folk

With their innate power of invisibility, pixies rarely appear unless they wish to be seen. In the Feywild and on the Material Plane, pixies etch patterns of frost on winter ponds and rouse the buds in springtime. They cause flowers to sparkle with summer dew, and color the leaves with the blazing hues of autumn.

## Pixie Dust

When pixies fly visibly, a shower of sparkling dust follows in their wake like the glittering tail of a shooting star. A mere sprinkle of pixie dust is said to be able to grant the power of flight, confuse a creature hopelessly, or send foes into a magical slumber.

Only pixies can use their dust to its full potential, but these fey are constantly sought out by mages and monsters seeking to study or master their power.

## Tiny Tricksters

While the arrival of visitors piques their curiosity, pixies are too shy to reveal themselves at first. They study the visitors from afar to gauge their temperament or play harmless tricks on them to measure their reactions. For example, pixies might tie a dwarf's boots together, create illusions of strange creatures or treasures, or use dancing lights to lead interlopers astray. If the visitors respond with hostility, the pixies give them a wide berth. If the visitors are good natured, the pixies are likely to be emboldened and more friendly. The fey might even emerge and offer to guide their "guests" along a safe route or invite them to a tiny yet satisfying feast prepared in their honor.

## Opposed to Violence

Unlike their fey cousins, the sprites, pixies abhor weapons and would sooner flee than get into a physical altercation with any enemy.

> [!quote]- A quote from Rivergleam, pixie fashionista  
> 
> Petal gowns and acorn caps are so last summer!


```statblock
"name": "Pixie"
"size": "Tiny"
"type": "fey"
"alignment": "Neutral Good"
"ac": !!int "15"
"hp": !!int "1"
"hit_dice": "1d4 - 1"
"stats":
- !!int "2"
- !!int "20"
- !!int "8"
- !!int "10"
- !!int "14"
- !!int "15"
"speed": "10 ft., fly 30 ft."
"skillsaves":
  "Stealth": !!int "7"
  "Perception": !!int "4"
"senses": "passive Perception 14"
"languages": "Sylvan"
"cr": "1/4"
"traits":
- "desc": "The pixie's innate spellcasting ability is Charisma (spell save DC 12).\
    \ It can innately cast the following spells, requiring only its pixie dust as\
    \ a component:\n\nAt will: [druidcraft](/3-Mechanics/CLI/spells/druidcraft.md)\n\
    \n1/day each: [confusion](/3-Mechanics/CLI/spells/confusion.md), [dancing\
    \ lights](/3-Mechanics/CLI/spells/dancing-lights.md), [detect evil and good](/3-Mechanics/CLI/spells/detect-evil-and-good.md),\
    \ [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [entangle](/3-Mechanics/CLI/spells/entangle.md), [fly](/3-Mechanics/CLI/spells/fly.md),\
    \ [phantasmal force](/3-Mechanics/CLI/spells/phantasmal-force.md), [polymorph](/3-Mechanics/CLI/spells/polymorph.md),\
    \ [sleep](/3-Mechanics/CLI/spells/sleep.md)"
  "name": "Innate Spellcasting"
- "desc": "The pixie has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The pixie magically turns [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ until its [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ ends (as if concentrating on a spell). Any equipment the pixie wears or carries\
    \ is [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible) with it."
  "name": "Superior Invisibility"
"source":
- "MM"
- "GoS"
- "ERLW"
- "IMR"
- "EGW"
- "WBtW"
- "KftGV"
- "SatO"
- "ToFW"
- "BMT"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Pixie.webp"
```
^statblock

```encounter-table
name: Pixie
creatures:
 - 1: Pixie
```

## Environment

forest