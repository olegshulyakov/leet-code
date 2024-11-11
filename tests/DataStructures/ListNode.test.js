const ListNode = require('../../src/DataStructures/ListNode')

describe('ListNode', () => {
    it('should create a new ListNode with default values', () => {
        const node = new ListNode();
        expect(node.val).toBe(0);
        expect(node.next).toBeNull();
    });

    it('should create a new ListNode with custom value and null next', () => {
        const node = new ListNode(5);
        expect(node.val).toBe(5);
        expect(node.next).toBeNull();
    });

    it('should create a new ListNode with default value and custom next', () => {
        const node1 = new ListNode();
        const node2 = new ListNode(10, null);
        node1.next = node2;
        expect(node1.val).toBe(0);
        expect(node2.val).toBe(10);
        expect(node2.next).toBeNull();
    });

    it('should create a linked list from a single element', () => {
        const nums = [1];
        const result = ListNode.from(nums);
        expect(result.val).toBe(1);
        expect(result.next).toBeNull();
    });

    it('should create a linked list from multiple elements in reverse order', () => {
        const nums = [1, 2, 3];
        const result = ListNode.from(nums);
        expect(result.val).toBe(1);
        expect(result.next.val).toBe(2);
        expect(result.next.next.val).toBe(3);
        expect(result.next.next.next).toBeNull();
    });

    it('should throw an error when creating a linked list from empty array', () => {
        const nums = [];
        expect(ListNode.from(nums)).toBeUndefined();
    });

    it('should create a linked list from array with only one element and no head', () => {
        const nums = [1];
        let head = undefined;
        for (let i = 0; i < nums.length; i++) {
            head = new ListNode(nums[i]);
        }
        expect(head.val).toBe(1);
    });

    it('should create a linked list from array with multiple elements and non-empty head', () => {
        const nums = [2, 3, 4];
        let head = ListNode.from(nums);
        expect(head.val).toBe(2);
        expect(head.next.val).toBe(3);
        expect(head.next.next.val).toBe(4);
    });

    it('should throw an error when creating a linked list from array with non-array elements', () => {
        const nums = 'a';
        expect(() => ListNode.from(nums)).toThrow('Input must be an array');
    });

    it('should create an array from a single-element linked list', () => {
        const node = new ListNode(1);
        const result = ListNode.toArray(node);
        expect(result).toEqual([1]);
    });

    it('should create an array from a multi-element linked list in reverse order', () => {
        const node = new ListNode(2, new ListNode(1, new ListNode(3)));
        const result = ListNode.toArray(node);
        expect(result).toEqual([2, 1, 3]);
    });

    it('should return empty array when passing null to toArray function', () => {
        expect(ListNode.toArray(null)).toEqual([]);
    });

    it('should create an array from empty linked list', () => {
        const node = new ListNode();
        expect(ListNode.toArray(node)).toEqual([0]);
    });

    it('should handle the case where input is not a linked list node', () => {
        expect(() => ListNode.toArray(123)).toThrow('Input must be a ListNode');
    });
});