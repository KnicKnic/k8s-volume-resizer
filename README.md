# Goal
To auto Resize a kubernetes volume if it starts to run out of space.

## Usage

This project is intended to be passed in the PVC you wish to control. The PVC will auto get resized with by an increase

## TODO

- [x] Create github action build script
- [x] create empty project
- [ ] Add ability to auto resize PVC
- [ ] Determine if StatefulSets can be auto resized
- [ ] Add e2e tests
- [ ] Add fake tests
- [ ] Add capability to detect what type of threshold triggers resize
    - [ ] percentage free space
    - [ ] amount free space
- [ ] add metrics for failure
- [ ] add warnings or triggers

