import unittest

import collector.document as document
from tests.util import Case


class TestIndex(unittest.TestCase):
    def test_parse_document(self):
        cases: list[Case] = [
            Case("emptry", "", ""),
        ]

        for c in cases:
            with self.subTest(c.name):
                result = document.parse_document(c.input)
                self.assertEqual(c.expected, result)


if __name__ == "__main__":
    unittest.main()
