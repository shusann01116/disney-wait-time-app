#!/bin/bash

(cd collector && python3 -m unittest discover ../tests/ -t ..)
