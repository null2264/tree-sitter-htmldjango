import XCTest
import SwiftTreeSitter
import TreeSitterHtmljinja

final class TreeSitterHtmljinjaTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_htmljinja())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Htmljinja grammar")
    }
}
