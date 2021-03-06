#!/usr/bin/env python3
# vi:nu:et:sts=4 ts=4 sw=4

''' Build the Application(s)

This module builds the go application generated by genapp.

The module must be executed from the repository that contains 'scripts'
directory.

'''


#   This is free and unencumbered software released into the public domain.
#
#   Anyone is free to copy, modify, publish, use, compile, sell, or
#   distribute this software, either in source code form or as a compiled
#   binary, for any purpose, commercial or non-commercial, and by any
#   means.
#
#   In jurisdictions that recognize copyright laws, the author or authors
#   of this software dedicate any and all copyright interest in the
#   software to the public domain. We make this dedication for the benefit
#   of the public at large and to the detriment of our heirs and
#   successors. We intend this dedication to be an overt act of
#   relinquishment in perpetuity of all present and future rights to this
#   software under copyright law.
#
#   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
#   EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
#   MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
#   IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
#   OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
#   ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE
#   OR OTHER DEALINGS IN THE SOFTWARE.
#
#   For more information, please refer to <http://unlicense.org/>

import sys
sys.path.insert(0, './scripts')
import util                             # pylint: disable=wrong-import-position


################################################################################
#                           Object Classes and Functions
################################################################################

#---------------------------------------------------------------------
#                   Main Command Execution Class
#---------------------------------------------------------------------

class Main(util.MainBase):
    """ Main Command Execution Class
    """

    def exec_pgm(self):
        """ Execute updating the main logic.
        """

        if self.args.verbose > 0:
            print('*****************************************')
            print('*       Building the Applications       *')
            print('*****************************************')
            print()

        result = self.do_cmd("./scripts/ci/build/build.py", 'app01ma')
        if result != 0:
            print("Error: Failed to build 'app01ma'!\n")
            return
        result = self.do_cmd("./scripts/ci/build/build.py", 'app01ms')
        if result != 0:
            print("Error: Failed to build 'app01ms'!\n")
            return
        result = self.do_cmd("./scripts/ci/build/build.py", 'app01my')
        if result != 0:
            print("Error: Failed to build 'app01my'!\n")
            return
        result = self.do_cmd("./scripts/ci/build/build.py", 'app01pg')
        if result != 0:
            print("Error: Failed to build 'app01pg'!\n")
            return
        result = self.do_cmd("./scripts/ci/build/build.py", 'app01sq')
        if result != 0:
            print("Error: Failed to build 'app01sq'!\n")
            return


################################################################################
#                           Command-line interface
################################################################################

if __name__ == '__main__':
    Main().run()
