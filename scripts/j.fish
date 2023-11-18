#!/usr/bin/env fish

# Copyright 2023 chaopeng@chaopeng.me
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# j is used to actually cd to the bookmarked dir.
function j
  set -l dir $(to find $argv[1])
  if test $status -eq 0
    cd $dir
  end
end

function __update_to_script --on-event fish_postexec
  if string match -qr 'to [save|add|delete|del|rm]' $argv[1]
    # saved bookmark is changed. source the new autocomplete script
    to genj fish > ~/.config/fish/completions/j.fish
    source ~/.config/fish/completions/j.fish
  end
end
